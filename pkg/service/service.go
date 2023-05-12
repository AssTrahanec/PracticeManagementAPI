package service

import (
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Student interface {
}
type Practice interface {
}

type Company interface {
}

type Request interface {
}

type User interface {
}

type University interface {
}
type Service struct {
	Authorization
	User
	Company
	Request
	Practice
	University
	Student
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
