package service

import (
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"
)

type PracticeService struct {
	repo repository.Practice
}

func NewPracticeService(repo repository.Practice) *PracticeService {
	return &PracticeService{repo: repo}
}

func (s *PracticeService) Create(userId int, practice model.Practice) (int, error) {
	return s.repo.Create(userId, practice)
}
func (s *PracticeService) GetAllOfCurrentUser(userId int) ([]model.Practice, error) {
	return s.repo.GetAllOfCurrentUser(userId)
}
func (s *PracticeService) GetAll(userId int) ([]model.Practice, error) {
	return s.repo.GetAll(userId)
}
func (s *PracticeService) GetAllPublic(userId int) ([]model.Practice, error) {
	return s.repo.GetAllPublic(userId)
}
func (s *PracticeService) GetById(userId, id int) (model.Practice, error) {
	return s.repo.GetById(userId, id)
}

func (s *PracticeService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
func (s *PracticeService) Update(userId, id int, input model.UpdatePracticeInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
