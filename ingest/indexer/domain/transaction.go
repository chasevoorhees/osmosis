package domain

import (
	"time"

	"github.com/cometbft/cometbft/abci/types"
)

type Transaction struct {
	Height             uint64        `json:"height"`
	BlockTime          time.Time     `json:"timestamp"`
	Sender             string        `json:"sender"`
	GasWanted          uint64        `json:"gas_wanted"`
	GasUsed            uint64        `json:"gas_used"`
	Fees               string        `json:"fees"`
	MessageType        string        `json:"msg_type"`
	TransactionHash    string        `json:"tx_hash"`
	TransactionIndexId int           `json:"tx_index_id"`
	Events             []types.Event `json:"events"`
	IngestedAt         time.Time     `json:"ingested_at"`
}
