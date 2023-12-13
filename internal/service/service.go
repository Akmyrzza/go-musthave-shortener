package service

import (
	"errors"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"math/rand"

	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
)

var RandLength = 16

type Repository interface {
	CreateShortURL(userID, shortURL, originalURL string) (string, error)
	GetOriginalURL(originalURL string) (string, error)
	PingStore() error
	CreateShortURLs(userID string, urls []model.ReqURL) ([]model.ReqURL, error)
	GetAllURLs(userID string) ([]model.ResURL, error)
}

type ServiceURL struct {
	Repository Repository
}

func NewServiceURL(r Repository) *ServiceURL {
	return &ServiceURL{
		Repository: r,
	}
}

func (s *ServiceURL) CreateShortURL(userID, originalURL string) (string, bool, error) {
	for {
		shortURL := randString()
		id, err := s.Repository.CreateShortURL(userID, originalURL, shortURL)
		if err != nil {
			if errors.Is(err, cerror.ErrAlreadyExist) {
				continue
			}
			return "", false, fmt.Errorf("creating id error: %w", err)
		}

		if id == "" {
			return shortURL, false, nil
		}

		return id, true, nil
	}
}

func (s *ServiceURL) GetOriginalURL(shortURL string) (string, error) {
	originalURL, ok := s.Repository.GetOriginalURL(shortURL)
	return originalURL, ok
}

func (s *ServiceURL) Ping() error {
	return s.Repository.PingStore()
}

func (s *ServiceURL) CreateShortURLs(userID string, urls []model.ReqURL) ([]model.ReqURL, error) {
	for i := range urls {
		shortURL := randString()
		urls[i].ShortURL = shortURL
	}

	res, err := s.Repository.CreateShortURLs(userID, urls)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ServiceURL) GetAllURLs(userID string) ([]model.ResURL, error) {
	data, err := s.Repository.GetAllURLs(userID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, RandLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
