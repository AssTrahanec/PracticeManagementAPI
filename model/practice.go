package model

type Practice struct {
	ID                    int    `json:"id"`
	CompanyID             int    `json:"company_id"`
	Duration              string `json:"duration"`
	StartDate             string `json:"start_date"`
	IsPayment             bool   `json:"is_payment"`
	PaymentAmount         string `json:"payment_amount"`
	WorkingHours          string `json:"working_hours"`
	Expectations          string `json:"expectations"`
	Benefits              string `json:"benefits"`
	Restrictions          string `json:"restrictions"`
	RequiredSkills        string `json:"required_skills"`
	AdditionalInformation string `json:"additional_information"`
}
