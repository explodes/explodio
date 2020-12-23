package main

import (
	"context"
	"flag"
	"github.com/explodes/explodio/stand"
	"github.com/explodes/explodio/stand/health"
	"google.golang.org/grpc"
	"os"
	"time"
)

var (
	addr    = flag.String("addr", "", "Health check address.")
	timeout = flag.Duration("timeout", 10*time.Second, "Run gRPC health check at the provided address.")
)

func main() {
	flag.Parse()

	logger := stand.NewStdoutLogger()
	if *addr == "" {
		logger.Errorf("no address specified for health check")
		os.Exit(1)
	}
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("did not connect: %v", err)
		os.Exit(1)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logger.Warnf("error closing connection: %v", err)
		}
	}()
	client := health.NewHealthClient(conn)

	ctx, cancelFunc := context.WithTimeout(context.Background(), *timeout)
	defer cancelFunc()
	_, err = client.Check(ctx, &health.HealthRequest{})
	if err != nil {
		logger.Errorf("health check failed: %v", err)
		os.Exit(1)
	}

	// Ok.
	os.Exit(0)
}
