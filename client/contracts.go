package main

import (
	"context"
	"fmt"
	"graph/pb"
	"log"
	"os"
	"strings"
	"time"
)

func createGraph(address, graph string){
	edges := make([]*pb.Edge, 0, 0)
	if graph != "null" {
		edgePairs := strings.Split(graph, ",")
		for i := 0; i < len(edgePairs); i++ {
			if !strings.Contains(edgePairs[i], "-") {
				fmt.Println("Invalid edge format", edgePairs[i])
				os.Exit(1)
			}
			edge := strings.Split(edgePairs[i], "-")
			if len(edge) != 2 {
				fmt.Println("Invalid edge format", edgePairs[i])
				os.Exit(1)
			}
			edges = append(edges, &pb.Edge{Source: edge[0], Dest: edge[1]})
		}
	}
	request := &pb.CreateRequest{Edges:edges}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	connector := NewConnector(address)
	conn := connector.OpenConn()
	defer conn.Close()

	client := pb.NewGraphServiceClient(conn)
	response, err := client.CreateGraph(ctx, request)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Println("GraphId:", response.GraphId)
}

func findPath(address, graphId, src, dest string){

	request := &pb.ShortestPathRequest{GraphId:graphId, Source: src, Destination: dest}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	connector := NewConnector(address)
	conn := connector.OpenConn()
	defer conn.Close()

	client := pb.NewGraphServiceClient(conn)
	response, err := client.ShortestPath(ctx, request)
	if err != nil {
		log.Fatalf("%v", err)
	}else {
		log.Println("Shortest Path:", response.ShortestPath)
	}
}

func deleteGraph(address, graphId string){

	request := &pb.DeleteRequest{GraphId:graphId}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	connector := NewConnector(address)
	conn := connector.OpenConn()
	defer conn.Close()

	client := pb.NewGraphServiceClient(conn)
	response, err := client.DeleteGraph(ctx, request)
	if err != nil {
		log.Fatalf("%v", err)
	}else {
		log.Println("Message:", response.Message)
	}
}
