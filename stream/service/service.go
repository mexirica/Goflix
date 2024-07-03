package service

import (
	"io"
	"os"
)

type VideoService interface {
	StreamVideo(filename string) (io.ReadCloser, error)
}

type videoService struct{}

func NewVideoService() VideoService {
	return &videoService{}
}

func (s *videoService) StreamVideo(filename string) (io.ReadCloser, error) {
	filePath := os.Getenv("MEDIA_PATH") + filename + ".mp4"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
