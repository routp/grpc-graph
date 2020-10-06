package model

import (
	"golang.org/x/net/context"
	"graph/pb"
)

type ShortestPathRequest struct {
	ID   string
	Src  string
	Dest string
}


type ShortestPathResponse struct {
	ShortestPath 	string `json:"shortestPath,omitempty"`
	Err     		string `json:json:"err,omitempty"`
}

func DecodePathRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ShortestPathRequest)
	return ShortestPathRequest{
		ID:   req.GraphId,
		Src:  req.Source,
		Dest: req.Destination,
	}, nil
}

func EncodePathResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(ShortestPathResponse)
	return &pb.ShortestPathResponse{
		ShortestPath: resp.ShortestPath,
		Err: resp.Err,
	}, nil
}