package service

import (
	"errors"
	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
	"math/rand"
)

type Repository interface {
	CreateID(id, originalURL string) error
	GetURL(id string) (string, bool)
}

type ServiceURL struct {
	Repository Repository
}

func NewServiceURL(r Repository) *ServiceURL {
	return &ServiceURL{
		Repository: r,
	}
}
func (s *ServiceURL) CreateID(originalURL string) string {
	for {
		id := randString()
		err := s.Repository.CreateID(id, originalURL)
		if err != nil {
			if errors.Is(err, cerror.ErrAlreadyExist) {
				continue
			} else {
				return ""
			}
		}

		return id
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
