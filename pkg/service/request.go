package service

import (
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"
)

type RequestService struct {
	repo repository.Request
}

func NewRequestService(repo repository.Request) *RequestService {
	return &RequestService{repo: repo}
}

func (s *RequestService) Create(userId int, request model.Request) (int, error) {
	return s.repo.Create(userId, request)
}
func (s *RequestService) GetAll(userId int) ([]model.Request, error) {
	return s.repo.GetAll(userId)
}
func (s *RequestService) GetById(userId, id int) (model.Request, error) {
	return s.repo.GetById(userId, id)
}

func (s *RequestService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
func (s *RequestService) Update(userId, id int, input model.UpdateRequestInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
