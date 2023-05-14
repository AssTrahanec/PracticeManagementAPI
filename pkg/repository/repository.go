package repository

import (
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/jmoiron/sqlx"
)

const (
	studentsTable     = "students"
	companiesTable    = "companies"
	requestsTable     = "requests"
	usersTable        = "users"
	practicesTable    = "practices"
	universitiesTable = "universities"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(login, password string) (model.User, error)
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
type Repository struct {
	Authorization
	User
	Company
	Request
	Practice
	University
	Student
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Request:       NewRequestPostgres(db),
		Company:       NewCompanyPostgres(db),
		Student:       NewStudentPostgres(db),
	}
}
