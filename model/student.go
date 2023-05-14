package model

import "errors"

type Student struct {
	ID           int      `json:"student_id" db:"student_id"`
	Specialty    *string  `json:"specialty" db:"specialty"`
	AvgMark      *float64 `json:"avg_mark" db:"avg_mark"`
	StudentName  *string  `json:"student_name" db:"student_name"`
	StudentEmail *string  `json:"student_email" db:"student_email"`
	StudentPhone *string  `json:"student_phone_number" db:"student_phone_number"`
	Speciality   *string  `json:"speciality" db:"speciality"`
}

type StudentResponse struct {
	Student
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UpdateStudentInput struct {
	Specialty    *string  `json:"specialty"`
	AvgMark      *float64 `json:"avg_mark"`
	StudentName  *string  `json:"student_name"`
	StudentEmail *string  `json:"student_email"`
	StudentPhone *string  `json:"student_phone_number"`
	Speciality   *string  `json:"speciality"`
}

func (i UpdateStudentInput) Validate() error {
	if i.Specialty == nil &&
		i.AvgMark == nil &&
		i.StudentName == nil &&
		i.StudentEmail == nil &&
		i.StudentPhone == nil &&
		i.Speciality == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
