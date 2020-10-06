package service

import (
	"context"
	"fmt"
	"graph/collections"
	"graph/gerror"
	"graph/logger"
	"graph/pb"
	"graph/store"
	"graph/traversal"
)

type GraphService interface {
	CreateGraph(ctx context.Context, edges []*pb.Edge) (string, error)
	ShortestPath(ctx context.Context, id, src, dest string) (string, error)
	DeleteGraph(ctx context.Context, id string) (string, error)
}

type GraphServiceImpl struct{}

// Create a graph
func (GraphServiceImpl) CreateGraph(ctx context.Context, edges []*pb.Edge) (string, error) {
	var graphId string
	var err error
	strEdges := transform(edges)
	graph, gErr := collections.ConstructGraph(strEdges)
	if gErr != nil {
		err = gErr.Error()
		logger.GetLogger().Error("Graph creation failed. error: %s", gErr.Message())
	}else{
		graphId = graph.GraphId
		store.Add(graphId, strEdges)
		logger.GetLogger().Info("Graph created with id %s", graphId)
	}
	return graphId, err
}

// Find shortest path
func (GraphServiceImpl) ShortestPath(ctx context.Context, id, src, dest string) (string, error) {
	var shortestPath string
	var err error

	edges := store.Get(id)
	if edges == nil || len(edges) == 0 {
		gErr := gerror.NewGraphError(gerror.IDNotFoundCode, fmt.Sprintf(gerror.IDNotFound, id))
		logger.GetLogger().Errorf("Shortest Path finding failed. error: %s", gErr.Message())
		err = gErr.Error()
	}else {
		graph, gErr := collections.ConstructGraphWithId(id, edges)
		if gErr != nil {
			err = gErr.Error()
			logger.GetLogger().Errorf("Shortest Path finding failed. error: %s", gErr.Message())
		} else {
			listPath := traversal.ShortestPath(graph, src, dest)
			shortestPath = traversal.StrPath(listPath)
			logger.GetLogger().Infof("Shortest Path found for graph %s from %s to %s: %s", id, src, dest, shortestPath)
		}
	}
	return shortestPath, err
}

// delete a graph
func (GraphServiceImpl) DeleteGraph(ctx context.Context, id string) (string, error){
	var message string
	var err error
	isDeleted := store.Remove(id)
	if isDeleted {
		message = fmt.Sprintf("Graph with id %s deleted", id)
		logger.GetLogger().Infom(message)
	}else{
		gErr := gerror.NewGraphError(gerror.IDNotFoundCode, fmt.Sprintf(gerror.IDNotFound, id))
		logger.GetLogger().Errorf("Delete graph failed. error: %s", gErr.Message())
		err = gErr.Error()
	}
	return message, err
}

func transform(edges []*pb.Edge) [][]string {
	var strGraph [][]string
	for i:=0; i < len(edges); i++ {
		strGraph = append(strGraph, []string{edges[i].Source, edges[i].Dest})
	}
	return strGraph
}