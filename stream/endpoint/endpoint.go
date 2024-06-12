package endpoint

import (
	"context"
	"io"

	"github.com/go-kit/kit/endpoint"
	"github.com/mexirica/goflix/stream/service"
)

// StreamVideoRequest holds the parameters for the StreamVideo endpoint
type StreamVideoRequest struct {
	Filename string
}

// StreamVideoResponse holds the response values for the StreamVideo endpoint
type StreamVideoResponse struct {
	File io.ReadCloser `json:"-"`
	Err  error         `json:"-"`
}

// MakeStreamVideoEndpoint creates the StreamVideo endpoint
func MakeStreamVideoEndpoint(svc service.VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StreamVideoRequest)
		file, err := svc.StreamVideo(req.Filename)
		return &StreamVideoResponse{File: file, Err: err}, nil
	}
}
