package repository

import (
	"fmt"
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type RequestPostgres struct {
	db *sqlx.DB
}

func NewRequestPostgres(db *sqlx.DB) *RequestPostgres {
	return &RequestPostgres{db: db}
}

func (r *RequestPostgres) Create(userId int, request model.Request) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (company_id, student_id, status, projects_link, language_skills, work_experience, skills) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", requestsTable)
	row := tx.QueryRow(createListQuery, request.CompanyID, userId, request.Status, request.ProjectsLink, request.LanguageSkills, request.WorkExperience, request.Skills)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
func (r *RequestPostgres) GetAll(userId int) ([]model.Request, error) {
	var requests []model.Request
	query := fmt.Sprintf("SELECT * FROM %s WHERE student_id = $1 OR company_id = $1;", requestsTable)
	err := r.db.Select(&requests, query, userId)

	return requests, err
}
func (r *RequestPostgres) GetById(userId, id int) (model.Request, error) {
	var request model.Request
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", requestsTable)
	err := r.db.Get(&request, query, id)

	return request, err
}
func (r *RequestPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = $1",
		requestsTable)
	_, err := r.db.Exec(query, id)

	return err
}
func (r *RequestPostgres) Update(userId, id int, input model.UpdateRequestInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Skills != nil {
		setValues = append(setValues, fmt.Sprintf("skills=$%d", argId))
		args = append(args, *input.Skills)
		argId++
	}
	if input.ProjectsLink != nil {
		setValues = append(setValues, fmt.Sprintf("projects_link=$%d", argId))
		args = append(args, *input.ProjectsLink)
		argId++
	}
	if input.LanguageSkills != nil {
		setValues = append(setValues, fmt.Sprintf("language_skills=$%d", argId))
		args = append(args, *input.LanguageSkills)
		argId++
	}
	if input.WorkExperience != nil {
		setValues = append(setValues, fmt.Sprintf("work_experience=$%d", argId))
		args = append(args, *input.WorkExperience)
		argId++
	}
	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", requestsTable, setQuery, argId)
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
