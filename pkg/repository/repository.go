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
	}
}
