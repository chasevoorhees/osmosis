package mocks

import (
	"context"
	"fmt"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/domain"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/domain/mvc"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/router/usecase/pools"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/router/usecase/route"
)

type PoolsUsecaseMock struct {
	Pools        []domain.PoolI
	TickModelMap map[uint64]domain.TickModel
}

// GetRoutesFromCandidates implements mvc.PoolsUsecase.
// Note that taker fee are ignored and not set
// Note that tick models are not set
func (r *PoolsUsecaseMock) GetRoutesFromCandidates(ctx context.Context, candidateRoutes route.CandidateRoutes, takerFeeMap domain.TakerFeeMap, tokenInDenom string, tokenOutDenom string) ([]route.RouteImpl, error) {
	finalRoutes := make([]route.RouteImpl, 0, len(candidateRoutes.Routes))
	for _, candidateRoute := range candidateRoutes.Routes {

		routablePools := make([]domain.RoutablePool, 0, len(candidateRoute.Pools))
		for _, candidatePool := range candidateRoute.Pools {
			// Get the pool data for routing
			var foundPool domain.PoolI
			for _, pool := range r.Pools {
				if pool.GetId() == candidatePool.ID {
					foundPool = pool
				}
			}

			if foundPool == nil {
				return nil, fmt.Errorf("pool with id %d not found in pools use case mock", candidatePool.ID)
			}

			// TODO: note that taker fee is force set to zero
			routablePool := pools.NewRoutablePool(foundPool, candidatePool.TokenOutDenom, osmomath.ZeroDec())
			routablePools = append(routablePools, routablePool)
		}

		finalRoutes = append(finalRoutes, route.RouteImpl{
			Pools: routablePools,
		})
	}

	return finalRoutes, nil
}

// GetAllPools implements domain.PoolsUsecase.
func (pm *PoolsUsecaseMock) GetAllPools(ctx context.Context) ([]domain.PoolI, error) {
	return pm.Pools, nil
}

// GetTickModelMap implements mvc.PoolsUsecase.
func (pm *PoolsUsecaseMock) GetTickModelMap(ctx context.Context, poolIDs []uint64) (map[uint64]domain.TickModel, error) {
	return pm.TickModelMap, nil
}

var _ mvc.PoolsUsecase = &PoolsUsecaseMock{}
