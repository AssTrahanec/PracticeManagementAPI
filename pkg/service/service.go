package service

import (
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Student interface {
	GetAll(userId int) ([]model.Student, error)
	GetById(userId, id int) (model.Student, error)
	Delete(userId, id int) error
	Update(userId, id int, input model.UpdateStudentInput) error
}
type Practice interface {
}

type Company interface {
	GetAll(userId int) ([]model.Company, error)
	GetById(userId, id int) (model.Company, error)
	Delete(userId, id int) error
	Update(userId, id int, input model.UpdateCompanyInput) error
}

type Request interface {
	Create(userId int, request model.Request) (int, error)
	GetAll(userId int) ([]model.Request, error)
	GetById(userId, id int) (model.Request, error)
	Delete(userId, id int) error
	Update(userId, id int, input model.UpdateRequestInput) error
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
		Request:       NewRequestService(repos.Request),
		Company:       NewCompanyService(repos.Company),
		Student:       NewStudentService(repos.Student),
	}
}
