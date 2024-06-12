package service

import (
	"io"
	"os"
)

// VideoService defines the video streaming service interface
type VideoService interface {
	StreamVideo(filename string) (io.ReadCloser, error)
}

// videoService is an implementation of VideoService
type videoService struct{}

func NewVideoService() VideoService {
	return &videoService{}
}

func (s *videoService) StreamVideo(filename string) (io.ReadCloser, error) {
	filePath := "videos/" + filename + ".mp4"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
