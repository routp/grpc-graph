package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"graph/service/model"
)

type GraphServiceEndpoints struct {
	CreateGraphEndpoint 	endpoint.Endpoint
	ShortestPathEndpoint 	endpoint.Endpoint
	DeleteGraphEndpoint 	endpoint.Endpoint
}

func MakeGraphEndpoints(graphSvc GraphService) GraphServiceEndpoints {
	return GraphServiceEndpoints{
		CreateGraphEndpoint: 	buildCreateGraphEndpoint(graphSvc),
		ShortestPathEndpoint: 	buildShortestPathEndpoint(graphSvc),
		DeleteGraphEndpoint: 	buildDeleteGraphEndpoint(graphSvc),
	}
}

func buildCreateGraphEndpoint(graphSvc GraphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.CreateRequest)
		graphId, err := graphSvc.CreateGraph(ctx, req.Edges)
		if err != nil {
			return nil, err
		}
		return model.CreateResponse{ID: graphId}, nil
	}
}

func buildShortestPathEndpoint(graphSvc GraphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.ShortestPathRequest)
		path,err := graphSvc.ShortestPath(ctx, req.ID, req.Src, req.Dest)
		if err != nil {
			return nil, err
		}
		return model.ShortestPathResponse{ShortestPath: path}, nil
	}
}

func buildDeleteGraphEndpoint(graphSvc GraphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.DeleteRequest)
		message, err := graphSvc.DeleteGraph(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return model.DeleteResponse{Message: message}, nil
	}
}
