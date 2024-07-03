package endpoint

import (
	"context"
	"io"

	"github.com/go-kit/kit/endpoint"
	"github.com/mexirica/goflix/stream/service"
)

type StreamVideoRequest struct {
	Filename string
}

type StreamVideoResponse struct {
	File io.ReadCloser `json:"-"`
	Err  error         `json:"-"`
}

func MakeStreamVideoEndpoint(svc service.VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StreamVideoRequest)
		file, err := svc.StreamVideo(req.Filename)
		return &StreamVideoResponse{File: file, Err: err}, nil
	}
}
