package transport

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/mexirica/goflix/stream/service"
	"io"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mexirica/goflix/stream/endpoint"
)

// decodeStreamVideoRequest decodes the request from HTTP to the endpoint format
func decodeStreamVideoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	filename := chi.URLParam(r, "filename")
	return endpoint.StreamVideoRequest{Filename: filename}, nil
}

// encodeStreamVideoResponse encodes the response from the endpoint to HTTP
func encodeStreamVideoResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(*endpoint.StreamVideoResponse)
	if resp.Err != nil {
		w.WriteHeader(http.StatusNotFound)
		return json.NewEncoder(w).Encode(map[string]string{"message": "error", "details": "File not found"})
	}

	defer resp.File.Close()
	w.Header().Set("Content-Type", "video/mp4")
	io.Copy(w, resp.File)
	return nil
}

// MakeHandler returns an HTTP handler that makes use of the endpoints
func MakeHandler(s service.VideoService) http.Handler {

	streamHandler := httptransport.NewServer(
		endpoint.MakeStreamVideoEndpoint(s),
		decodeStreamVideoRequest,
		encodeStreamVideoResponse,
	)

	r := chi.NewRouter()
	r.Get("/{filename}", streamHandler.ServeHTTP)

	return r
}
