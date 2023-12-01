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

type InMemory struct {
	dataURL map[string]string
}

type LocalRepository struct {
	file         *os.File
	maxRecord    int
	InMemoryRepo *InMemory
}

type tmpStorage struct {
	ID          string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

var Permission = 0600

func NewRepo(filePath string) (*LocalRepository, error) {
	inMemory := new(InMemory)
	inMemory.dataURL = make(map[string]string)

	if filePath == "" {
		return &LocalRepository{
			file:         nil,
			maxRecord:    0,
			InMemoryRepo: inMemory,
		}, nil
	}

	local, err := initFileDatabase(filePath, inMemory)
	if err != nil {
		return nil, err
	}

	return local, nil
}

func initFileDatabase(filePath string, InMemory *InMemory) (*LocalRepository, error) {
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

		InMemory.dataURL[tmpRecord.ShortURL] = tmpRecord.OriginalURL
		maxRecord, err = strconv.Atoi(tmpRecord.ID)
		if err != nil {
			return nil, cerror.ErrStringToInt
		}
	}

	ptrLocal := new(LocalRepository)
	ptrLocal.file = fileDB
	ptrLocal.maxRecord = maxRecord
	ptrLocal.InMemoryRepo = InMemory
	return ptrLocal, nil
}

func (s *LocalRepository) CreateID(shortURL, originalURL string) error {
	_, found := s.InMemoryRepo.dataURL[shortURL]
	if found {
		return cerror.ErrAlreadyExist
	}

	s.InMemoryRepo.dataURL[shortURL] = originalURL

	if s.file == nil {
		return nil
	}

	if err := saveInLocalDatabase(s, shortURL, originalURL); err != nil {
		return err
	}

	return nil
}

func saveInLocalDatabase(s *LocalRepository, shortURL, originalURL string) error {
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

func (s *LocalRepository) GetURL(id string) (string, bool) {
	originalURL, ok := s.InMemoryRepo.dataURL[id]
	return originalURL, ok
}
