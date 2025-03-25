package service

import "social-network-api/pkg/repository"

type Authorization interface{}

type Posts interface{}

type Users interface{}

type Service struct {
	Authorization
	Posts
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
