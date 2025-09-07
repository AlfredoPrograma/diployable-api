package services

import (
	"context"

	"github.com/AlfredoPrograma/diployable/db"
)

type UsersService interface {
	GetUserByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error)
}

type usersService struct {
	db *db.Queries
}

func (s usersService) GetUserByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error) {
	return s.db.GetUserByEmail(ctx, email)
}

func NewUsersService(db *db.Queries) UsersService {
	return usersService{
		db: db,
	}
}
