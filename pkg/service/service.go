package service

import "github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
