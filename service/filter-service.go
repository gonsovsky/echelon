package service

import (
	"context"
	"log"
)

type FilterServiceGrpcImpl struct {
}

func NewFilterServiceGrpcImpl() *FilterServiceGrpcImpl {
	return &FilterServiceGrpcImpl{}
}

func (serviceImpl *FilterServiceGrpcImpl) Add(ctx context.Context, in *Record) (*AddResponse, error) {
	log.Printf("Received request: %+v\n", in)

	return &AddResponse{
		Result: true,
		Record: in,
	}, nil
}
