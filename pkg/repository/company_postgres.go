package repository

import (
	"fmt"
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

type CompanyPostgres struct {
	db *sqlx.DB
}

func NewCompanyPostgres(db *sqlx.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (r *CompanyPostgres) GetAll(userId int) ([]model.Company, error) {
	var companies []model.Company
	query := fmt.Sprintf("SELECT * FROM %s", companiesTable)
	err := r.db.Select(&companies, query)

	return companies, err
}

func (r *CompanyPostgres) GetById(userId, id int) (model.Company, error) {
	var company model.Company
	query := fmt.Sprintf("SELECT * FROM %s WHERE company_id = $1", companiesTable)
	err := r.db.Get(&company, query, id)

	return company, err
}

func (r *CompanyPostgres) Delete(userId, id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteCompanyQuery := fmt.Sprintf("DELETE FROM %s WHERE company_id = $1", companiesTable)
	_, err = tx.Exec(deleteCompanyQuery, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	deleteUserQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err = tx.Exec(deleteUserQuery, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *CompanyPostgres) Update(userId, id int, input model.UpdateCompanyInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}
	if input.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argID))
		args = append(args, *input.Address)
		argID++
	}
	if input.ContactPerson != nil {
		setValues = append(setValues, fmt.Sprintf("contact_person=$%d", argID))
		args = append(args, *input.ContactPerson)
		argID++
	}
	if input.PhoneNumber != nil {
		setValues = append(setValues, fmt.Sprintf("phone_number=$%d", argID))
		args = append(args, *input.PhoneNumber)
		argID++
	}
	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argID))
		args = append(args, *input.Email)
		argID++
	}
	if input.Website != nil {
		setValues = append(setValues, fmt.Sprintf("website=$%d", argID))
		args = append(args, *input.Website)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE company_id = $%d", companiesTable, setQuery, argID)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)

	return err
}
