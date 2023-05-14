package model

import "errors"

type Request struct {
	ID             int    `json:"id" db:"id"`
	CompanyID      int    `json:"company_id" db:"company_id"`
	StudentID      int    `json:"student_id" db:"student_id"`
	Status         string `json:"status" db:"status"`
	ProjectsLink   string `json:"projects_link" db:"projects_link"`
	LanguageSkills string `json:"language_skills" db:"language_skills"`
	WorkExperience int    `json:"work_experience" db:"work_experience"`
	Skills         string `json:"skills" db:"skills"`
}

type UpdateRequestInput struct {
	Status         *string `json:"status"`
	ProjectsLink   *string `json:"projects_link"`
	LanguageSkills *string `json:"language_skills"`
	WorkExperience *int    `json:"work_experience"`
	Skills         *string `json:"skills"`
}

func (i UpdateRequestInput) Validate() error {
	if i.Status == nil &&
		i.ProjectsLink == nil &&
		i.LanguageSkills == nil &&
		i.WorkExperience == nil &&
		i.Skills == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
