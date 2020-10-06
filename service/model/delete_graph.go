package model

import (
	"golang.org/x/net/context"
	"graph/pb"
)

type DeleteRequest struct {
	ID string
}


type DeleteResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:""err,omitempty"`
}

func DecodeDeleteGraphRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.DeleteRequest)
	return DeleteRequest{
		ID: req.GraphId,
	}, nil
}

func EncodeDeleteGraphResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(DeleteResponse)
	return &pb.DeleteResponse{
		Message: resp.Message,
		Err: resp.Err,
	}, nil
}