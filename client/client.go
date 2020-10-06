package main

import (
	"google.golang.org/grpc"
	"log"
)

type Connector struct {
	address string
}

func (c *Connector) OpenConn() *grpc.ClientConn {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	return conn
}

func NewConnector(address string) *Connector {
	return &Connector{address}
}
