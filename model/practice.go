package model

import (
	"errors"
	"time"
)

type Practice struct {
	ID                    int        `json:"id" db:"id"`
	CompanyID             int        `json:"company_id" db:"company_id"`
	Duration              *string    `json:"duration" db:"duration"`
	StartDate             *time.Time `json:"start_date" db:"start_date"`
	IsPayment             *bool      `json:"is_payment" db:"is_payment"`
	PaymentAmount         *string    `json:"payment_amount" db:"payment_amount"`
	WorkingHours          *string    `json:"working_hours" db:"working_hours"`
	Expectations          *string    `json:"expectations" db:"expectations"`
	Benefits              *string    `json:"benefits" db:"benefits"`
	Restrictions          *string    `json:"restrictions" db:"restrictions"`
	RequiredSkills        *string    `json:"required_skills" db:"required_skills"`
	AdditionalInformation *string    `json:"additional_information" db:"additional_information"`
	PracticeStatus        *string    `json:"practice_status" db:"practice_status"`
}

type UpdatePracticeInput struct {
	Duration              *string    `json:"duration"`
	StartDate             *time.Time `json:"start_date"`
	IsPayment             *bool      `json:"is_payment"`
	PaymentAmount         *string    `json:"payment_amount"`
	WorkingHours          *string    `json:"working_hours"`
	Expectations          *string    `json:"expectations"`
	Benefits              *string    `json:"benefits"`
	Restrictions          *string    `json:"restrictions"`
	RequiredSkills        *string    `json:"required_skills"`
	AdditionalInformation *string    `json:"additional_information"`
	PracticeStatus        *string    `json:"practice_status"`
}

func (i UpdatePracticeInput) Validate() error {
	if i.Duration == nil &&
		i.StartDate == nil &&
		i.IsPayment == nil &&
		i.PaymentAmount == nil &&
		i.WorkingHours == nil &&
		i.Expectations == nil &&
		i.Benefits == nil &&
		i.Restrictions == nil &&
		i.RequiredSkills == nil &&
		i.AdditionalInformation == nil &&
		i.PracticeStatus == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
