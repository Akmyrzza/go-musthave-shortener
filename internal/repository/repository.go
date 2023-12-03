package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
)

type Repository interface {
	CreateID(id, originalURL string) error
	GetURL(id string) (string, bool)
}

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

func NewRepo(filePath string) (Repository, error) {

	inMemoryStore := initInMemoryStore()

	if filePath == "" {
		return inMemoryStore, nil
	}

	fileStore, err := initFileStore(inMemoryStore, filePath)
	if err != nil {
		return nil, fmt.Errorf("initialiizing file store: %w", err)
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

func (s *inMemory) CreateID(shortURL, originalURL string) error {
	_, found := s.dataURL[shortURL]
	if found {
		return cerror.ErrAlreadyExist
	}

	s.dataURL[shortURL] = originalURL

	return nil
}

func (s *inMemory) GetURL(id string) (string, bool) {
	originalURL, ok := s.dataURL[id]
	return originalURL, ok
}

func (s *localRepository) CreateID(shortURL, originalURL string) error {
	if err := s.inMemoryRepo.CreateID(shortURL, originalURL); err != nil {
		return err
	}

	if err := saveInLocalDatabase(s, shortURL, originalURL); err != nil {
		return err
	}

	return nil
}

func (s *localRepository) GetURL(id string) (string, bool) {
	return s.inMemoryRepo.GetURL(id)
}

func saveInLocalDatabase(s *localRepository, shortURL, originalURL string) error {
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
