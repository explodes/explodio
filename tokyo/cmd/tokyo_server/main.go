package main

import (
	"fmt"
	"github.com/explodes/explodio/stand"
	"github.com/explodes/explodio/tokyo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	logger := stand.NewStdoutLogger()
	storage, err := tokyo.ConnectStorage()
	if err != nil {
		panic(fmt.Errorf("unable to connect to storage: %v", err))
	}
	server, err := tokyo.NewTokyoServer(logger, storage)
	if err != nil {
		panic(fmt.Errorf("unable to create server: %v", err))
	}

	lis, err := net.Listen("tcp", stand.RequireEnv("PB_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	tokyo.RegisterTokyoServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
