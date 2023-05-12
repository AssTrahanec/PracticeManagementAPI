package model

type Company struct {
	ID            int    `json:"company_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Address       string `json:"address"`
	ContactPerson string `json:"contact_person"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	Website       string `json:"website"`
}

type CompanyResponse struct {
	Company
	Username string `json:"username"`
	Role     string `json:"role"`
}
