package stand

import (
	"context"
	"flag"
	"github.com/explodes/explodio/stand"
	"google.golang.org/grpc"
	"os"
	"time"
)

var (
	performCheck = flag.Bool("healthCheck", false, "Run gRPC health check at a provided address.")
	addr         = flag.String("healthAddr", "", "Health check address.")
	timeout      = flag.Duration("healthTimeout", 10*time.Second, "Run gRPC health check at the provided address.")
)

type healthServer struct {
	UnimplementedHealthServer
}

var _ HealthServer = (*healthServer)(nil)

func AttachHealthServer(s *grpc.Server) {
	server := &healthServer{}
	RegisterHealthServer(s, server)
}

func ReplyToHealthCheck() {
	logger := stand.NewStdoutLogger()
	flag.Parse()
	if !(*performCheck) {
		return
	}
	if *addr == "" {
		logger.Errorf("no address specified for health check")
		os.Exit(1)
	}
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		logger.Errorf("did not connect: %w", err)
		os.Exit(1)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logger.Warnf("error closing connection: %w", err)
		}
	}()
	client := NewHealthClient(conn)

	ctx, cancelFunc := context.WithTimeout(context.Background(), *timeout)
	defer cancelFunc()
	_, err = client.Check(ctx, &HealthRequest{})
	if err != nil {
		logger.Errorf("health check failed: %w", err)
		os.Exit(1)
	}

	// Ok.
	os.Exit(0)
}

func (h healthServer) Check(ctx context.Context, request *HealthRequest) (*HealthResponse, error) {
	return &HealthResponse{}, nil
}
