package gerror

import (
	"github.com/pkg/errors"

)

const (
	EdgeCountError = "edges count must be two. index %d"
	EdgeCountCode  = "G001"
	IDNotFound     = "Graph id %s not found"
	IDNotFoundCode = "G002"
)

type GraphError struct {
	errorCode string
	errorMsg  string
	error     error
}

func NewGraphError(code, message string) *GraphError {
	return &GraphError{code, message, errors.New(message)}
}

func (ge *GraphError) Error() error {
	return ge.error
}

func (ge *GraphError) Message() string {
	return ge.errorMsg
}

func (ge *GraphError) Code() string {
	return ge.errorCode
}
