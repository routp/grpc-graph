package service

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"graph/pb"
	"graph/service/model"
)

type gRpcServer struct {
	create 	grpctransport.Handler
	path 	grpctransport.Handler
	delete	grpctransport.Handler
}

func (s *gRpcServer) CreateGraph(ctx context.Context, r *pb.CreateRequest) (*pb.CreateResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateResponse), nil
}

func (s *gRpcServer) ShortestPath(ctx context.Context, sp *pb.ShortestPathRequest) (*pb.ShortestPathResponse, error) {
	_, resp, err := s.path.ServeGRPC(ctx, sp)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ShortestPathResponse), nil
}

func (s *gRpcServer) DeleteGraph(ctx context.Context, d *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	_, resp, err := s.delete.ServeGRPC(ctx, d)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeleteResponse), nil
}

func NewGRPCServer(ctx context.Context, graphSvc GraphService) pb.GraphServiceServer {
	endpoints :=  MakeGraphEndpoints(graphSvc)
	return &gRpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateGraphEndpoint,
			model.DecodeCreateGraphRequest,
			model.EncodeCreateGraphResponse,
		),
		path: grpctransport.NewServer(
			endpoints.ShortestPathEndpoint,
			model.DecodePathRequest,
			model.EncodePathResponse,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteGraphEndpoint,
			model.DecodeDeleteGraphRequest,
			model.EncodeDeleteGraphResponse,
		),
	}
}