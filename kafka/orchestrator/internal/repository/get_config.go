package repository

import (
	"context"
	"orchestrator/internal/domain"
)

var findByTypeAndService = `---
select
	*
from
	route_config
where
	order_type = $1 and order_service = $2
limit 
	1;
`

func (repo *OrchestratorConfigImpl) GetConfig(ctx context.Context) (config domain.RouteConfig, err error) {
	findByTypeAndServiceStmt, err := repo.db.PrepareContext(ctx, findByTypeAndService)
	if err != nil {
		return domain.RouteConfig{}, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.RouteConfig{}, err
	}

	orchestratorType := ctx.Value(domain.Key("type")).(string)
	orchestratorService := ctx.Value(domain.Key("service")).(string)

	row := tx.StmtContext(ctx, findByTypeAndServiceStmt).QueryRowContext(ctx, orchestratorType, orchestratorService)
	if row.Err() != nil {
		return domain.RouteConfig{}, err
	}

	err = row.Scan(
		&config.OrderType,
		&config.OrderService,
		&config.TargetTopic,
	)

	if err != nil {
		return domain.RouteConfig{}, err
	}

	return
}
