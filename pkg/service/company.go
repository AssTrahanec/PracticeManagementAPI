package service

import (
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"
)

type CompanyService struct {
	repo repository.Company
}

func NewCompanyService(repo repository.Company) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) GetAll(userId int) ([]model.Company, error) {
	return s.repo.GetAll(userId)
}
func (s *CompanyService) GetById(userId, id int) (model.Company, error) {
	return s.repo.GetById(userId, id)
}

func (s *CompanyService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
func (s *CompanyService) Update(userId, id int, input model.UpdateCompanyInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
