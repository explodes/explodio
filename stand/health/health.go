package health

import (
	"context"
	"google.golang.org/grpc"
)

type healthServer struct {
	UnimplementedHealthServer
}

var _ HealthServer = (*healthServer)(nil)

func AttachHealthServer(s *grpc.Server) {
	server := &healthServer{}
	RegisterHealthServer(s, server)
}

func (h healthServer) Check(ctx context.Context, request *HealthRequest) (*HealthResponse, error) {
	return &HealthResponse{}, nil
}
