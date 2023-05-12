package repository

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
type Repository struct {
	Authorization
	User
	Company
	Request
	Practice
	University
	Student
}

func NewRepository() *Repository {
	return &Repository{}
}
