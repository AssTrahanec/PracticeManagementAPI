package repository

import (
	"fmt"
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, role) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := tx.QueryRow(query, user.Login, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	if user.Role == "student" {
		studentQuery := fmt.Sprintf("INSERT INTO %s (student_id) VALUES ($1)", studentsTable)
		_, err := tx.Exec(studentQuery, id)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	} else if user.Role == "company" {
		companyQuery := fmt.Sprintf("INSERT INTO %s (company_id) VALUES ($1)", companiesTable)
		_, err := tx.Exec(companyQuery, id)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}
func (r *AuthPostgres) GetUser(login, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
