package services

import (
	"context"

	"github.com/AlfredoPrograma/diployable/db"
	"github.com/AlfredoPrograma/diployable/dtos"
)

type UsersService interface {
	GetUserByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error)
	CreateUser(ctx context.Context, data dtos.CreateUserDTO) error
}

type usersService struct {
	db             *db.Queries
	encryptService EncryptService
}

func (s usersService) CreateUser(ctx context.Context, data dtos.CreateUserDTO) error {
	hash, err := s.encryptService.Encrypt(data.Password)

	if err != nil {
		return err
	}

	user := db.CreateUserParams{
		Email:    data.Email,
		Password: hash,
	}

	return s.db.CreateUser(ctx, user)
}

func (s usersService) GetUserByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error) {
	return s.db.GetUserByEmail(ctx, email)
}

func NewUsersService(db *db.Queries, encryptService EncryptService) UsersService {
	return usersService{
		db:             db,
		encryptService: encryptService,
	}
}
