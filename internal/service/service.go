package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"math/rand"

	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
)

var RandLength = 16

type Repository interface {
	CreateShortURL(shortURL, originalURL string) (string, error)
	GetOriginalURL(originalURL string) (string, error)
	Ping(ctx context.Context) error
	CreateShortURLs(urls []model.ReqURL) ([]model.ReqURL, error)
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
		id, err := s.Repository.CreateShortURL(originalURL, shortURL)
		if err != nil {
			if errors.Is(err, cerror.ErrAlreadyExist) {
				return id, cerror.ErrAlreadyExist
			}
			return "", fmt.Errorf("creating id error: %w", err)
		}

		return id, nil
	}
}

func (s *ServiceURL) GetOriginalURL(shortURL string) (string, error) {
	originalURL, ok := s.Repository.GetOriginalURL(shortURL)
	return originalURL, ok
}

func (s *ServiceURL) Ping(ctx context.Context) error {
	return s.Repository.Ping(ctx)
}

func (s *ServiceURL) CreateShortURLs(urls []model.ReqURL) ([]model.ReqURL, error) {
	for i := range urls {
		shortURL := randString()
		urls[i].ShortURL = shortURL
	}

	res, err := s.Repository.CreateShortURLs(urls)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, RandLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
