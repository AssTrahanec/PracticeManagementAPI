package model

import "errors"

type Company struct {
	ID            int     `json:"company_id" db:"company_id"`
	Name          *string `json:"name" db:"name"`
	Description   *string `json:"description" db:"description"`
	Address       *string `json:"address" db:"address"`
	ContactPerson *string `json:"contact_person" db:"contact_person"`
	PhoneNumber   *string `json:"phone_number" db:"phone_number"`
	Email         *string `json:"email" db:"email"`
	Website       *string `json:"website" db:"website"`
}

type UpdateCompanyInput struct {
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Address       *string `json:"address"`
	ContactPerson *string `json:"contact_person"`
	PhoneNumber   *string `json:"phone_number"`
	Email         *string `json:"email"`
	Website       *string `json:"website"`
}

type CompanyResponse struct {
	Company
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (i UpdateCompanyInput) Validate() error {
	if i.Name == nil &&
		i.Description == nil &&
		i.Address == nil &&
		i.ContactPerson == nil &&
		i.PhoneNumber == nil &&
		i.Email == nil &&
		i.Website == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
