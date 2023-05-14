package service

import (
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"
)

type StudentService struct {
	repo repository.Student
}

func NewStudentService(repo repository.Student) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) GetAll(userId int) ([]model.Student, error) {
	return s.repo.GetAll(userId)
}
func (s *StudentService) GetById(userId, id int) (model.Student, error) {
	return s.repo.GetById(userId, id)
}

func (s *StudentService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
func (s *StudentService) Update(userId, id int, input model.UpdateStudentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
