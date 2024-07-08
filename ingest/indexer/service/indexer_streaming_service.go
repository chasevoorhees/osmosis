package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	// wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	gammtypes "github.com/osmosis-labs/osmosis/v25/x/gamm/types"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v25/x/poolmanager/types"

	"github.com/osmosis-labs/osmosis/v25/ingest/indexer/domain"
)

var _ baseapp.StreamingService = (*indexerStreamingService)(nil)

// ind is a streaming service that processes block data and ingests it into the indexer
type indexerStreamingService struct {
	writeListeners map[storetypes.StoreKey][]storetypes.WriteListener

	// manages tracking of whether the node is code started
	coldStartManager domain.ColdStartManager

	client domain.Publisher

	keepers domain.Keepers

	txDecoder sdk.TxDecoder

	txnIndexId int
}

// New creates a new sqsStreamingService.
// writeListeners is a map of store keys to write listeners.
// sqsIngester is an ingester that ingests the block data into SQS.
// poolTracker is a tracker that tracks the pools that were changed in the block.
// nodeStatusChecker is a checker that checks if the node is syncing.
func New(writeListeners map[storetypes.StoreKey][]storetypes.WriteListener, coldStartManager domain.ColdStartManager, client domain.Publisher, keepers domain.Keepers, txDecoder sdk.TxDecoder) baseapp.StreamingService {
	return &indexerStreamingService{

		writeListeners: writeListeners,

		coldStartManager: coldStartManager,

		client: client,

		keepers: keepers,

		txDecoder: txDecoder,
	}
}

// Close implements baseapp.StreamingService.
func (s *indexerStreamingService) Close() error {
	return nil
}

// ListenBeginBlock implements baseapp.StreamingService.
func (s *indexerStreamingService) ListenBeginBlock(ctx context.Context, req types.RequestBeginBlock, res types.ResponseBeginBlock) error {
	s.txnIndexId++
	return nil
}

// ListenCommit implements baseapp.StreamingService.
func (s *indexerStreamingService) ListenCommit(ctx context.Context, res types.ResponseCommit) error {
	return nil
}

// ListenDeliverTx implements baseapp.StreamingService.
func (s *indexerStreamingService) ListenDeliverTx(ctx context.Context, req types.RequestDeliverTx, res types.ResponseDeliverTx) error {
	// Publish the transaction data
	err := s.publishTxn(ctx, req, res)
	if err != nil {
		return err
	}
	return nil
}

// publishBlock publishes the block data to the indexer.
func (s *indexerStreamingService) publishBlock(ctx context.Context, req types.RequestEndBlock) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	height := (uint64)(req.GetHeight())
	timeEndBlock := sdkCtx.BlockTime().UTC()
	chainId := sdkCtx.ChainID()
	gasConsumed := sdkCtx.GasMeter().GasConsumed()
	block := domain.Block{
		ChainId:     chainId,
		Height:      height,
		BlockTime:   timeEndBlock,
		GasConsumed: gasConsumed,
	}
	return s.client.PublishBlock(sdkCtx, block)
}

// publishTxn publishes the transaction data to the indexer.
func (s *indexerStreamingService) publishTxn(ctx context.Context, req types.RequestDeliverTx, res types.ResponseDeliverTx) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// Decode the transaction
	tx, err := s.txDecoder(req.GetTx())
	if err != nil {
		return err
	}

	// Calculate the transaction hash
	txHash := strings.ToUpper(hex.EncodeToString(tmhash.Sum(req.GetTx())))

	// Gas data
	gasWanted := res.GasWanted
	gasUsed := res.GasUsed

	// TO DO - Fees. Where is it?
	fees := "0"

	// Iterate through the messages in the transaction
	var msgType string
	var sender string
	txMessages := tx.GetMsgs()
	for _, txMsgGeneric := range txMessages {
		//
		// Other msg types to consider:
		//   import transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
		//   import wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
		//
		// transfertypes.MsgTransfer
		// wasmtypes.MsgExecuteContract
		//
		if txMsg, ok := txMsgGeneric.(*poolmanagertypes.MsgSwapExactAmountIn); ok {
			// txMsg is a pointer to poolmanagertypes.MsgSwapExactAmountIn
			// Look for token_swapped event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			tokenIn := txMsg.GetTokenIn()
			tokenOutMinAmount := txMsg.TokenOutMinAmount
			fmt.Printf("sender: %s, tokenIn: %s, tokenOutMinAmount: %s", sender, tokenIn, tokenOutMinAmount)
			routes := txMsg.GetRoutes()
			for _, route := range routes {
				poolId := route.GetPoolId()
				tokenOutDenom := route.GetTokenOutDenom()
				fmt.Printf("poolId: %d, tokenOutDenom: %s", poolId, tokenOutDenom)
			}
		}
		if txMsg, ok := txMsgGeneric.(*poolmanagertypes.MsgSwapExactAmountOut); ok {
			// txMsg is a pointer to poolmanagertypes.MsgSwapExactAmountOut
			// Look for token_swapped event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			tokenOut := txMsg.GetTokenOut()
			tokenInMaxAmount := txMsg.TokenInMaxAmount
			fmt.Printf("sender: %s, tokenIn: %s, tokenInMaxAmount: %s", sender, tokenOut, tokenInMaxAmount)
			routes := txMsg.GetRoutes()
			for _, route := range routes {
				poolId := route.GetPoolId()
				tokenInDenom := route.GetTokenInDenom()
				fmt.Printf("poolId: %d, tokenInDenom: %s", poolId, tokenInDenom)
			}
		}
		if txMsg, ok := txMsgGeneric.(*poolmanagertypes.MsgSplitRouteSwapExactAmountIn); ok {
			// txMsg is a pointer to poolmanagertypes.MsgSplitRouteSwapExactAmountIn
			// Look for token_swapped event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			tokenInDenom := txMsg.GetTokenInDenom()
			tokenOutMinAmount := txMsg.TokenOutMinAmount
			fmt.Printf("sender: %s, tokenInDenom: %s, tokenOutMinAmount: %s", sender, tokenInDenom, tokenOutMinAmount)
			routes := txMsg.GetRoutes()
			for _, route := range routes {
				tokenInAmount := route.TokenInAmount
				pools := route.GetPools()
				for _, pool := range pools {
					poolId := pool.GetPoolId()
					tokenOutDenom := pool.GetTokenOutDenom()
					fmt.Printf("poolId: %d, tokenOutDenom: %s, tokenInAmount: %s", poolId, tokenOutDenom, tokenInAmount)
				}
			}
		}
		if txMsg, ok := txMsgGeneric.(*poolmanagertypes.MsgSplitRouteSwapExactAmountOut); ok {
			// txMsg is a pointer to poolmanagertypes.MsgSplitRouteSwapExactAmountOut
			// Look for token_swapped event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			tokenOutDenom := txMsg.GetTokenOutDenom()
			tokenInMaxAmount := txMsg.TokenInMaxAmount
			fmt.Printf("sender: %s, tokenIn: %s, tokenInMaxAmount: %s", sender, tokenOutDenom, tokenInMaxAmount)
			routes := txMsg.GetRoutes()
			for _, route := range routes {
				fmt.Println("route", route.GetPools())
				pools := route.GetPools()
				for _, pool := range pools {
					fmt.Println("poolId", pool.GetPoolId())
					fmt.Println("tokenOutDenom", pool.GetTokenInDenom())
				}
			}
		}
		if txMsg, ok := txMsgGeneric.(*gammtypes.MsgJoinPool); ok {
			// txMsg is a pointer to gammtypes.MsgJoinPool
			// Look for pool_joined event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			poolId := txMsg.GetPoolId()
			shareOutAmount := txMsg.ShareOutAmount
			fmt.Printf("sender: %s, poolId: %s, shareOutAmount: %s", sender, strconv.FormatUint(poolId, 10), shareOutAmount)
			tokenInMaxs := txMsg.GetTokenInMaxs()
			for _, tokenInMax := range tokenInMaxs {
				denom := tokenInMax.GetDenom()
				amount := tokenInMax.Amount
				fmt.Printf("denom: %s, amount: %s", denom, amount)
			}
		}
		if txMsg, ok := txMsgGeneric.(*gammtypes.MsgJoinSwapExternAmountIn); ok {
			// txMsg is a pointer to gammtypes.MsgJoinSwapExternAmountIn
			// Look for pool_joined event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			poolId := txMsg.GetPoolId()
			tokenIn := txMsg.GetTokenIn()
			shareOutMinAmount := txMsg.ShareOutMinAmount
			fmt.Printf("sender: %s, poolId: %s, tokenIn: %s, shareOutMinAmount: %s", sender, strconv.FormatUint(poolId, 10), tokenIn, shareOutMinAmount)
		}
		if txMsg, ok := txMsgGeneric.(*gammtypes.MsgJoinSwapShareAmountOut); ok {
			// txMsg is a pointer to gammtypes.MsgJoinSwapExternAmountOut
			// Look for pool_joined event in events
			msgType = txMsg.Type()
			poolId := txMsg.GetPoolId()
			sender = txMsg.GetSender()
			tokenInDenom := txMsg.GetTokenInDenom()
			tokenInMaxAmount := txMsg.TokenInMaxAmount
			fmt.Printf("sender: %s, poolId: %s, tokenInDenom: %s, tokenInMaxAmount: %s", sender, strconv.FormatUint(poolId, 10), tokenInDenom, tokenInMaxAmount)
		}
		if txMsg, ok := txMsgGeneric.(*gammtypes.MsgExitPool); ok {
			// txMsg is a pointer to gammtypes.MsgExitPool
			// Look for pool_exited event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			poolId := txMsg.GetPoolId()
			shareOutAmount := txMsg.ShareInAmount
			fmt.Printf("sender: %s, poolId: %s, shareOutAmount: %s", sender, strconv.FormatUint(poolId, 10), shareOutAmount)
			tokenInMins := txMsg.GetTokenOutMins()
			for _, tokenOutMin := range tokenInMins {
				denom := tokenOutMin.GetDenom()
				amount := tokenOutMin.Amount
				fmt.Printf("denom: %s, amount: %s", denom, amount)
			}
		}
		if txMsg, ok := txMsgGeneric.(*gammtypes.MsgExitSwapExternAmountOut); ok {
			// txMsg is a pointer to gammtypes.MsgExitSwapExternAmountOut
			// Look for pool_exited event in events
			msgType = txMsg.Type()
			pooldId := txMsg.GetPoolId()
			sender = txMsg.GetSender()
			tokenOut := txMsg.GetTokenOut()
			shareInMaxAmount := txMsg.ShareInMaxAmount
			fmt.Printf("sender: %s, poolId: %s, tokenOut: %s, shareInMaxAmount: %s", sender, strconv.FormatUint(pooldId, 10), tokenOut, shareInMaxAmount)
		}
		if txMsg, ok := txMsgGeneric.(*gammtypes.MsgExitSwapShareAmountIn); ok {
			// txMsg is a pointer to gammtypes.MsgExitSwapShareAmountIn
			// Look for pool_exited event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			poolId := txMsg.GetPoolId()
			tokenOutDenom := txMsg.TokenOutDenom
			tokenOutMinAmount := txMsg.TokenOutMinAmount
			shareInAmount := txMsg.ShareInAmount
			fmt.Printf("sender: %s, poolId: %s, tokenOutDenom: %s, tokenOutMinAmount: %s, shareInAmount: %s", sender, strconv.FormatUint(poolId, 10), tokenOutDenom, tokenOutMinAmount, shareInAmount)
		}
		if txMsg, ok := txMsgGeneric.(*gammtypes.MsgSwapExactAmountIn); ok {
			// txMsg is a pointer to gammtypes.MsgSwapExactAmountIn
			// Look for token_swapped event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			tokenIn := txMsg.GetTokenIn()
			tokenOutMinAmount := txMsg.TokenOutMinAmount
			fmt.Printf("sender: %s, tokenIn: %s, tokenOutMinAmount: %s", sender, tokenIn, tokenOutMinAmount)
			routes := txMsg.GetRoutes()
			for _, route := range routes {
				poolId := route.GetPoolId()
				tokenOutDenom := route.GetTokenOutDenom()
				fmt.Printf("poolId: %d, tokenOutDenom: %s", poolId, tokenOutDenom)
			}
		}
		if txMsg, ok := txMsgGeneric.(*gammtypes.MsgSwapExactAmountOut); ok {
			// txMsg is a pointer to gammtypes.MsgSwapExactAmountOut
			// Look for token_swapped event in events
			msgType = txMsg.Type()
			sender = txMsg.GetSender()
			tokenOut := txMsg.GetTokenOut()
			tokenInMaxAmount := txMsg.TokenInMaxAmount
			fmt.Printf("sender: %s, tokenIn: %s, tokenOutMinAmount: %s", sender, tokenOut, tokenInMaxAmount)
			routes := txMsg.GetRoutes()
			for _, route := range routes {
				poolId := route.GetPoolId()
				tokenInDenom := route.GetTokenInDenom()
				fmt.Printf("poolId: %d, tokenInDenom: %s", poolId, tokenInDenom)
			}
		}
	}

	if msgType == "" {
		// TO DO - should log this error
		return nil
	}

	events := res.GetEvents()
	txn := domain.Transaction{
		Height:          uint64(sdkCtx.BlockHeight()),
		BlockTime:       sdkCtx.BlockTime().UTC(),
		Sender:          sender,
		GasWanted:       uint64(gasWanted),
		GasUsed:         uint64(gasUsed),
		Fees:            fees,
		MessageType:     msgType,
		TransactionHash: txHash,
		Events:          events,
	}
	return s.client.PublishTransaction(sdkCtx, txn)
}

// ListenEndBlock implements baseapp.StreamingService.
func (s *indexerStreamingService) ListenEndBlock(ctx context.Context, req types.RequestEndBlock, res types.ResponseEndBlock) error {
	defer func() {
		s.txnIndexId = 0
	}()
	// Publish the block data
	err := s.publishBlock(ctx, req)
	if err != nil {
		return err
	}

	/*
		// If did not ingest initial data yet, ingest it now
		if !s.coldStartManager.HasIngestedInitialData() {
			sdkCtx := sdk.UnwrapSDKContext(ctx)

			var err error

			// Ingest the initial data
			s.keepers.BankKeeper.IterateTotalSupply(sdkCtx, func(coin sdk.Coin) bool {
				// Check if the denom should be filtered out and skip it if so
				if domain.ShouldFilterDenom(coin.Denom) {
					return false
				}

				// Publish the token supply
				err = s.client.PublishTokenSupply(sdkCtx, domain.TokenSupply{
					Denom:  coin.Denom,
					Supply: coin.Amount,
				})

				// Skip any error silently but log it.
				if err != nil {
					// TODO: alert
					sdkCtx.Logger().Error("failed to publish token supply", "error", err)
				}

				supplyOffset := s.keepers.BankKeeper.GetSupplyOffset(sdkCtx, coin.Denom)

				// If supply offset is non-zero, publish it.
				if !supplyOffset.IsZero() {
					// Publish the token supply offset
					err = s.client.PublishTokenSupplyOffset(sdkCtx, domain.TokenSupplyOffset{
						Denom:        coin.Denom,
						SupplyOffset: supplyOffset,
					})
				}

				return false
			})

			// Mark that the initial data has been ingested
			s.coldStartManager.MarkInitialDataIngested()
		}
	*/

	return nil
}

// Listeners implements baseapp.StreamingService.
func (s *indexerStreamingService) Listeners() map[storetypes.StoreKey][]storetypes.WriteListener {
	return s.writeListeners
}

// Stream implements baseapp.StreamingService.
func (s *indexerStreamingService) Stream(wg *sync.WaitGroup) error {
	return nil
}
