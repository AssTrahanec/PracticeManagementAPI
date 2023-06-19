package repository

import (
	"fmt"
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

type StudentPostgres struct {
	db *sqlx.DB
}

func NewStudentPostgres(db *sqlx.DB) *StudentPostgres {
	return &StudentPostgres{db: db}
}

func (r *StudentPostgres) GetAll(userId int) ([]model.Student, error) {
	var students []model.Student
	query := fmt.Sprintf("SELECT * FROM %s", studentsTable)
	err := r.db.Select(&students, query)

	return students, err
}

func (r *StudentPostgres) GetById(userId, id int) (model.Student, error) {
	var student model.Student
	query := "SELECT * FROM students WHERE student_id = $1"
	err := r.db.Get(&student, query, id)
	if err != nil {
		return model.Student{}, err
	}
	return student, nil
}

func (r *StudentPostgres) Create(student model.Student) (int, error) {
	query := "INSERT INTO students (specialty, avg_mark, student_name, student_email, student_phone_number, speciality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING student_id"
	row := r.db.QueryRow(query, student.Specialty, student.AvgMark, student.StudentName, student.StudentEmail, student.StudentPhone, student.Speciality)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *StudentPostgres) Update(userId, id int, input model.UpdateStudentInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Specialty != nil {
		setValues = append(setValues, fmt.Sprintf("specialty=$%d", argID))
		args = append(args, *input.Specialty)
		argID++
	}
	if input.AvgMark != nil {
		setValues = append(setValues, fmt.Sprintf("avg_mark=$%d", argID))
		args = append(args, *input.AvgMark)
		argID++
	}
	if input.StudentName != nil {
		setValues = append(setValues, fmt.Sprintf("student_name=$%d", argID))
		args = append(args, *input.StudentName)
		argID++
	}
	if input.StudentEmail != nil {
		setValues = append(setValues, fmt.Sprintf("student_email=$%d", argID))
		args = append(args, *input.StudentEmail)
		argID++
	}
	if input.StudentPhone != nil {
		setValues = append(setValues, fmt.Sprintf("student_phone_number=$%d", argID))
		args = append(args, *input.StudentPhone)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE students SET %s WHERE student_id = $%d", setQuery, argID)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *StudentPostgres) Delete(userId, id int) error {
	query := "DELETE FROM students WHERE student_id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
