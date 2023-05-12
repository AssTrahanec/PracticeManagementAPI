package repository

import "github.com/jmoiron/sqlx"

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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
