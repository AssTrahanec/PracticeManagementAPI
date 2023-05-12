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
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, role) values ($1, $2, $3) returning id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
