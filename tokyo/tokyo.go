package tokyo

import (
	"context"
	"github.com/explodes/explodio/stand"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type tokyoServer struct {
	UnimplementedTokyoServer
	logger  stand.Logger
	storage Storage
}

var _ TokyoServer = (*tokyoServer)(nil)

func NewTokyoServer(logger stand.Logger, storage Storage) (TokyoServer, error) {
	return &tokyoServer{
		logger:  logger,
		storage: storage,
	}, nil
}

func (t tokyoServer) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, error) {
	existing, err := t.storage.GetUser(ctx, request.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot lookup user")
	} else if existing != nil {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	if len(request.Password) < 8 {
		return nil, status.Error(codes.InvalidArgument, "password too short")
	}
	if !isValidEmail(request.Email) {
		return nil, status.Error(codes.InvalidArgument, "invalid email address")
	}

	hashed, err := HashPassword(request.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot hash password")
	}

	user, err := t.storage.CreateUser(ctx, request.Email, hashed)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot create user")
	}

	return &CreateResponse{
		User: &User{
			Uuid:    user.UUID,
			Email: request.Email,
		},
	}, nil
}
