package repository

import (
	"fmt"
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type PracticePostgres struct {
	db *sqlx.DB
}

func NewPracticePostgres(db *sqlx.DB) *PracticePostgres {
	return &PracticePostgres{db: db}
}

func (r *PracticePostgres) Create(userId int, practice model.Practice) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (company_id, duration, start_date, is_payment, payment_amount, working_hours, expectations, benefits, restrictions, required_skills, additional_information) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", practicesTable)
	row := tx.QueryRow(createListQuery, userId, practice.Duration, practice.StartDate, practice.IsPayment, practice.PaymentAmount, practice.WorkingHours, practice.Expectations, practice.Benefits, practice.Restrictions, practice.RequiredSkills, practice.AdditionalInformation)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
func (r *PracticePostgres) GetAllOfCurrentUser(userId int) ([]model.Practice, error) {
	var practices []model.Practice
	query := fmt.Sprintf("SELECT * FROM %s WHERE company_id = $1;", practicesTable)
	err := r.db.Select(&practices, query, userId)

	return practices, err
}
func (r *PracticePostgres) GetAll(userId int) ([]model.Practice, error) {
	var practices []model.Practice
	query := fmt.Sprintf("SELECT * FROM %s", practicesTable)
	err := r.db.Select(&practices, query)

	return practices, err
}
func (r *PracticePostgres) GetAllPublic(userId int) ([]model.Practice, error) {
	var practices []model.Practice
	query := fmt.Sprintf("SELECT * FROM %s WHERE practice_status = '%s'", practicesTable, "public")
	err := r.db.Select(&practices, query)
	return practices, err
}
func (r *PracticePostgres) GetById(userId, id int) (model.Practice, error) {
	var practice model.Practice
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", practicesTable)
	err := r.db.Get(&practice, query, id)

	return practice, err
}
func (r *PracticePostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = $1",
		practicesTable)
	_, err := r.db.Exec(query, id)

	return err
}
func (r *PracticePostgres) Update(userId, id int, input model.UpdatePracticeInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Duration != nil {
		setValues = append(setValues, fmt.Sprintf("duration=$%d", argId))
		args = append(args, input.Duration)
		argId++
	}
	if input.StartDate != nil {
		setValues = append(setValues, fmt.Sprintf("start_date=$%d", argId))
		args = append(args, input.StartDate)
		argId++
	}
	if input.IsPayment != nil {
		setValues = append(setValues, fmt.Sprintf("is_payment=$%d", argId))
		args = append(args, input.IsPayment)
		argId++
	}
	if input.PaymentAmount != nil {
		setValues = append(setValues, fmt.Sprintf("payment_amount=$%d", argId))
		args = append(args, input.PaymentAmount)
		argId++
	}
	if input.WorkingHours != nil {
		setValues = append(setValues, fmt.Sprintf("working_hours=$%d", argId))
		args = append(args, input.WorkingHours)
		argId++
	}
	if input.Expectations != nil {
		setValues = append(setValues, fmt.Sprintf("expectations=$%d", argId))
		args = append(args, input.Expectations)
		argId++
	}
	if input.Benefits != nil {
		setValues = append(setValues, fmt.Sprintf("benefits=$%d", argId))
		args = append(args, input.Benefits)
		argId++
	}
	if input.Restrictions != nil {
		setValues = append(setValues, fmt.Sprintf("restrictions=$%d", argId))
		args = append(args, input.Restrictions)
		argId++
	}
	if input.RequiredSkills != nil {
		setValues = append(setValues, fmt.Sprintf("required_skills=$%d", argId))
		args = append(args, input.RequiredSkills)
		argId++
	}
	if input.AdditionalInformation != nil {
		setValues = append(setValues, fmt.Sprintf("additional_information=$%d", argId))
		args = append(args, input.AdditionalInformation)
		argId++
	}
	if input.PracticeStatus != nil {
		setValues = append(setValues, fmt.Sprintf("practice_status=$%d", argId))
		args = append(args, input.PracticeStatus)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", practicesTable, setQuery, argId)
	args = append(args, id)
	logrus.SetLevel(logrus.DebugLevel)

	logger := logrus.StandardLogger() // Получение стандартного логгера
	// Вывод отладочной информации
	logger.Debugf("This is a debug message.")
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
