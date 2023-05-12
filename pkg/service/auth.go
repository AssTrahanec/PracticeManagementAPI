package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Asstrahanec/PracticeManagementAPI/model"
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/repository"
	"time"
)

const (
	salt       = "djskddjks313kjj1j3"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
