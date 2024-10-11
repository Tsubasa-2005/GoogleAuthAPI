package handler

import (
	"context"

	"GoogleAuthAPI/pkg/infra/rdb"
	"GoogleAuthAPI/ui/api"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	dbConn *pgxpool.Pool
	repo   *rdb.Queries
}

func (h *Handler) GetUsers(ctx context.Context) (*api.GetUsersOK, error) {
	//TODO implement me
	panic("implement me")
}

type SecurityHandler struct{}

func NewHandler(dbConn *pgxpool.Pool) *Handler {
	return &Handler{
		dbConn: dbConn,
		repo:   rdb.New(dbConn),
	}
}

func NewSecurityHandler() *SecurityHandler {
	return &SecurityHandler{}
}
