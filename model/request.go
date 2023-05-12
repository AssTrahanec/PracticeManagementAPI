package model

type Request struct {
	ID             int    `json:"id"`
	CompanyID      int    `json:"company_id"`
	StudentID      int    `json:"student_id"`
	Status         string `json:"status"`
	ProjectsLink   string `json:"projects_link"`
	LanguageSkills string `json:"language_skills"`
	WorkExperience int    `json:"work_experience"`
	Skills         string `json:"skills"`
}
