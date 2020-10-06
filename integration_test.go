package main_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"graph/logger"
	"graph/pb"
	"graph/service"
	"log"
	"net"
	"os"
	"testing"
)

const (
	bufSize = 1024 * 1024
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

func setup() {
	logConf := logger.Conf{}
	logger.InitLogger(logConf)
}

var listener *bufconn.Listener

func init() {
	var gSvc service.GraphService
	gSvc = service.GraphServiceImpl{}
	gSvcHandler := service.NewGRPCServer(context.Background(), gSvc)
	listener = bufconn.Listen(bufSize)

	server := grpc.NewServer()
	pb.RegisterGraphServiceServer(server, gSvcHandler)
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}
func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func TestIntegration(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "mockServer", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial mockServer: %v", err)
	}
	defer conn.Close()

	client := pb.NewGraphServiceClient(conn)

	// Create Graph
	resp, err := client.CreateGraph(ctx, &pb.CreateRequest{Edges:populateGraph()})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GraphId)
	t.Log("GraphId: ", resp.GraphId)

	// Find Shortest Path
	sp, err := client.ShortestPath(ctx, &pb.ShortestPathRequest{GraphId:resp.GraphId, Source: "D", Destination: "A"})
	assert.Nil(t, err)
	assert.NotNil(t, sp)
	assert.NotEmpty(t, sp.ShortestPath)
	t.Log("Shortest Path: ", sp.ShortestPath)
	assert.Equal(t,  "D -> B -> A", sp.ShortestPath)

	// Invalid graph id - shortest path
	sp, err = client.ShortestPath(ctx, &pb.ShortestPathRequest{GraphId:"a123", Source: "D", Destination: "A"})
	assert.Error(t, err)
	t.Log(err)

	// Delete graph
	dr, err := client.DeleteGraph(ctx, &pb.DeleteRequest{GraphId:resp.GraphId})
	assert.Nil(t, err)
	assert.NotNil(t, dr)
	assert.NotEmpty(t, dr.Message)

	// Invalid graph id - delete graph
	dr, err = client.DeleteGraph(ctx, &pb.DeleteRequest{GraphId:"a123"})
	assert.Error(t, err)
	t.Log(err)
}


func populateGraph() []*pb.Edge {
	edges := make([]*pb.Edge, 0, 20)
	edges= append(edges, &pb.Edge{Source: "A", Dest: "B"})
	edges= append(edges, &pb.Edge{Source: "A", Dest: "C"})
	edges= append(edges, &pb.Edge{Source: "B", Dest: "D"})
	return edges
}