package tokyo

import (
	"context"
	"database/sql"
	"github.com/explodes/explodio/stand"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type UserEntity struct {
	Id             int64
	UUID           string
	Email          string
	HashedPassword string
}

type Storage interface {
	GetUser(ctx context.Context, email string) (*UserEntity, error)
	CreateUser(ctx context.Context, email, hashedPassword string) (*UserEntity, error)
}

func ConnectStorage() (Storage, error) {
	db := stand.ConnectPostgres()
	return &storage{db: db}, nil
}

type storage struct {
	db *pg.DB
}

var _ Storage = (*storage)(nil)

func (s *storage) GetUser(ctx context.Context, email string) (*UserEntity, error) {
	user := new(UserEntity)
	err := s.db.Model(user).Context(ctx).Where("email = ?", email).Select()
	if err == sql.ErrNoRows {
		err = nil
	}
	return user, err
}

func (s *storage) CreateUser(ctx context.Context, email, hashedPassword string) (*UserEntity, error) {
	user := &UserEntity{
		UUID:           uuid.New().String(),
		Email:          email,
		HashedPassword: hashedPassword,
	}
	_, err := s.db.Model(user).Context(ctx).Insert()
	return user, err
}
