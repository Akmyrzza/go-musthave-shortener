package service

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/repository"
	"math/rand"
)

type ServiceURL struct {
	Repository repository.Repository
}

func NewServiceURL(r repository.Repository) *ServiceURL {
	return &ServiceURL{
		Repository: r,
	}
}
func (s *ServiceURL) CreateID(originalURL string) string {
	id := ""
	for {
		id = randString()
		_, ok := s.Repository.GetURL(id)
		if !ok {
			s.Repository.CreateID(id, originalURL)
			break
		}
	}

	return id
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
