package users

import (
	"context"

	"github.com/AlfredoPrograma/diployable/db"
)

type Service interface {
	GetUserByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error)
}

type service struct {
	db *db.Queries
}

func (s service) GetUserByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error) {
	return s.db.GetUserByEmail(ctx, email)
}

func NewService(db *db.Queries) Service {
	return service{
		db: db,
	}
}
