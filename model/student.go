package model

type Student struct {
	ID           int    `json:"student_id"`
	Specialty    string `json:"specialty"`
	AvgMark      int    `json:"avg_mark"`
	StudentName  string `json:"student_name"`
	StudentEmail string `json:"student_email"`
	StudentPhone string `json:"student_phone_number"`
	Speciality   string `json:"speciality"`
}

type StudentResponse struct {
	Student
	Username string `json:"username"`
	Role     string `json:"role"`
}
