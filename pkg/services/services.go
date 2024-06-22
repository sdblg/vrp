package services

import (
	"github.com/sdblg/vrp/pkg/configs"
)

type IService interface {
	Do()
}

func New(cfg configs.Config) (IService, error) {
	return &service{cfg: cfg}, nil
}

type service struct {
	cfg configs.Config
}
