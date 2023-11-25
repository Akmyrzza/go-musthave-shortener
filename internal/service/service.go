package service

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/repository"
	"math/rand"
)

type Service interface {
	CreateID(originalURL string) string
	GetURL(id string) (string, bool)
}

type ServiceURL struct {
	Repository repository.Repository
}

func NewServiceURL(r repository.Repository) *ServiceURL {
	return &ServiceURL{
		Repository: r,
	}
}
func (s *ServiceURL) CreateID(originalURL string) string {
	for {
		id := randString()
		found := s.Repository.CreateID(id, originalURL)
		if found == nil {
			return id
		}
	}
}

func (s *ServiceURL) GetURL(id string) (string, bool) {
	originalURL, ok := s.Repository.GetURL(id)
	return originalURL, ok
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
