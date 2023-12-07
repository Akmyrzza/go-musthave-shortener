package service

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
)

var RandLength = 8

type Repository interface {
	CreateShortURL(shortURL, originalURL string) error
	GetOriginalURL(originalURL string) (string, error)
	PingStore() error
}

type ServiceURL struct {
	Repository Repository
}

func NewServiceURL(r Repository) *ServiceURL {
	return &ServiceURL{
		Repository: r,
	}
}

func (s *ServiceURL) CreateShortURL(originalURL string) (string, error) {
	for {
		shortURL := randString()
		err := s.Repository.CreateShortURL(shortURL, originalURL)
		if err != nil {
			if errors.Is(err, cerror.ErrAlreadyExist) {
				continue
			}
			return "", fmt.Errorf("creating id error: %w", err)
		}

		return shortURL, nil
	}
}

func (s *ServiceURL) GetOriginalURL(shortURL string) (string, error) {
	originalURL, ok := s.Repository.GetOriginalURL(shortURL)
	return originalURL, ok
}

func (s *ServiceURL) Ping() error {
	return s.Repository.PingStore()
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, RandLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
