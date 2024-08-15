package repository

import (
	"context"
	"database/sql"
	"orchestrator/internal/domain"
)

type OrchestratorConfig interface {
	GetConfig(context.Context) (domain.RouteConfig, error)
}

type OrchestratorConfigImpl struct {
	db *sql.DB
}

func NewOrchestratorConfig(db *sql.DB) OrchestratorConfig {
	return &OrchestratorConfigImpl{db: db}
}
