package service

import (
	"GINTEST/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Findall() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(newvideo entity.Video) entity.Video {
	service.videos = append(service.videos, newvideo)
	return newvideo
}

func (service *videoService) Findall() []entity.Video {
	return service.videos
}
