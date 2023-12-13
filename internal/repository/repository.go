package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"io"
	"log"
	"os"
	"strconv"
)

type inMemory struct {
	dataURL map[string]string
}

type localRepository struct {
	file         *os.File
	maxRecord    int
	inMemoryRepo *inMemory
}

type tmpStorage struct {
	ID          string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

var Permission = 0600

func NewRepo(filePath string) (service.Repository, error) {
	inMemoryStore := initInMemoryStore()

	if filePath == "" {
		return inMemoryStore, nil
	}

	fileStore, err := initFileStore(inMemoryStore, filePath)
	if err != nil {
		return nil, fmt.Errorf("initialiizing file pgsql: %w", err)
	}

	return fileStore, nil
}

func initInMemoryStore() *inMemory {
	inMemory := new(inMemory)
	inMemory.dataURL = make(map[string]string)

	return inMemory
}

func initFileStore(m *inMemory, filePath string) (*localRepository, error) {
	fileDB, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(Permission))
	if err != nil {
		return nil, cerror.ErrOpenFileRepo
	}

	newDecoder := json.NewDecoder(fileDB)

	maxRecord := 0

	for {
		var tmpRecord tmpStorage

		if err := newDecoder.Decode(&tmpRecord); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("decoding error: %w", err)
		}

		m.dataURL[tmpRecord.ShortURL] = tmpRecord.OriginalURL
		maxRecord, err = strconv.Atoi(tmpRecord.ID)
		if err != nil {
			return nil, cerror.ErrStringToInt
		}
	}

	ptrLocal := new(localRepository)
	ptrLocal.file = fileDB
	ptrLocal.maxRecord = maxRecord
	ptrLocal.inMemoryRepo = m
	return ptrLocal, nil
}

func (s *inMemory) CreateShortURL(originalURL, shortURL string) (string, error) {
	_, found := s.dataURL[shortURL]
	if found {
		return "", cerror.ErrAlreadyExist
	}

	s.dataURL[shortURL] = originalURL

	return "", nil
}

func (s *inMemory) GetOriginalURL(id string) (string, error) {
	originalURL, ok := s.dataURL[id]
	if !ok {
		return "", fmt.Errorf("not found in base")
	}

	return originalURL, nil
}

func (s *inMemory) PingStore() error {
	return errors.New("no ping")
}

func (s *inMemory) CreateShortURLs(urls []model.ReqURL) ([]model.ReqURL, error) {
	for _, v := range urls {
		_, found := s.dataURL[v.ShortURL]
		if found {
			return nil, cerror.ErrAlreadyExist
		}

		s.dataURL[v.ShortURL] = v.OriginalURL
	}

	return urls, nil
}

func (s *inMemory) GetAllURLs() ([]model.ResURL, error) {
	var data []model.ResURL

	for key, value := range s.dataURL {
		var row model.ResURL

		row.ShortURL = key
		row.OriginalURL = value
		data = append(data, row)
	}

	return data, nil
}

func (s *localRepository) CreateShortURL(originalURL, shortURL string) (string, error) {
	id, err := s.inMemoryRepo.CreateShortURL(originalURL, shortURL)
	if err != nil {
		return id, err
	}

	if err := saveInLocalDatabase(s, originalURL, shortURL); err != nil {
		return "", err
	}

	return "", nil
}

func (s *localRepository) GetOriginalURL(originalURL string) (string, error) {
	return s.inMemoryRepo.GetOriginalURL(originalURL)
}

func (s *localRepository) PingStore() error {
	return errors.New("no ping")
}

func (s *localRepository) CreateShortURLs(urls []model.ReqURL) ([]model.ReqURL, error) {
	urls, err := s.inMemoryRepo.CreateShortURLs(urls)
	if err != nil {
		return nil, fmt.Errorf("saving in memory error: %w", err)
	}

	for i, v := range urls {
		if err := saveInLocalDatabase(s, v.OriginalURL, v.ShortURL); err != nil {
			return nil, fmt.Errorf("saving in file error: %w", err)
		}
		urls[i].OriginalURL = ""
	}

	return urls, nil
}

func (s *localRepository) GetAllURLs() ([]model.ResURL, error) {
	data, err := s.inMemoryRepo.GetAllURLs()
	return data, err
}

func saveInLocalDatabase(s *localRepository, originalURL, shortURL string) error {
	s.maxRecord++

	var tmpRecord tmpStorage
	tmpRecord.ID = strconv.Itoa(s.maxRecord)
	tmpRecord.ShortURL = shortURL
	tmpRecord.OriginalURL = originalURL

	data, err := json.Marshal(&tmpRecord)
	if err != nil {
		log.Fatalf("error: reading from json: %d", err)
	}

	data = append(data, '\n')
	_, err = s.file.Write(data)
	if err != nil {
		log.Fatalf("error: writing to json file: %d", err)
	}

	return nil
}
