package service

import (
	"errors"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/repository"
	"math/rand"

	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
)

var RandLength = 8

type Store interface {
	PingStore() error
}

type ServiceURL struct {
	Repository repository.Repository
	DB         Store
}

func NewServiceURL(r repository.Repository, db Store) *ServiceURL {
	return &ServiceURL{
		Repository: r,
		DB:         db,
	}
}

func (s *ServiceURL) CreateID(originalURL string) (string, error) {
	for {
		id := randString()
		err := s.Repository.CreateID(id, originalURL)
		if err != nil {
			if errors.Is(err, cerror.ErrAlreadyExist) {
				continue
			}
			return "", fmt.Errorf("creating id error: %w", err)
		}

		return id, nil
	}
}

func (s *ServiceURL) GetURL(id string) (string, bool) {
	originalURL, ok := s.Repository.GetURL(id)
	return originalURL, ok
}

func (s *ServiceURL) Ping() error {
	return s.DB.PingStore()
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, RandLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
