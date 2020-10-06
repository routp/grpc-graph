package model


import (
	"context"
	"graph/pb"
)

//request
type CreateRequest struct {
	Edges []*pb.Edge
}

//response
type CreateResponse struct {
	ID 		string `json:"graphId"`
	Err     string `json:"err,omitempty"`
}

func DecodeCreateGraphRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateRequest)
	return CreateRequest{
		Edges: req.Edges,
	}, nil
}

func EncodeCreateGraphResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(CreateResponse)
	return &pb.CreateResponse{
		GraphId: resp.ID,
		Err: resp.Err,
	}, nil
}