package repository

import (
	"encoding/json"
	"errors"
	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
	"io"
	"log"
	"os"
	"strconv"
)

type InMemory struct {
	dataURL map[string]string
	local   *LocalRepository
}

type LocalRepository struct {
	file      *os.File
	maxRecord int
}

type tmpStorage struct {
	ID          string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewInMemory(filePath string) (*InMemory, error) {
	if filePath == "" {
		return &InMemory{
			dataURL: make(map[string]string),
			local:   nil,
		}, nil
	}

	dataURL := make(map[string]string)

	local, err := initFileDatabase(filePath, &dataURL)
	if err != nil {
		return nil, err
	}

	return &InMemory{
		dataURL: dataURL,
		local:   local,
	}, nil
}

func initFileDatabase(filePath string, dataURL *map[string]string) (*LocalRepository, error) {

	fileDB, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	newDecoder := json.NewDecoder(fileDB)

	maxRecord := 0

	for {
		var tmpRecord tmpStorage

		if err := newDecoder.Decode(&tmpRecord); err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, err
			}
		}

		(*dataURL)[tmpRecord.ShortURL] = tmpRecord.OriginalURL
		maxRecord, err = strconv.Atoi(tmpRecord.ID)
		if err != nil {
			return nil, err
		}
	}

	ptrLocal := new(LocalRepository)
	ptrLocal.file = fileDB
	ptrLocal.maxRecord = maxRecord
	return ptrLocal, nil
}

func (s *InMemory) CreateID(shortURL, originalURL string) error {
	_, found := s.dataURL[shortURL]
	if found {
		return cerror.ErrAlreadyExist
	}

	s.dataURL[shortURL] = originalURL

	if s.local == nil {
		return nil
	}

	if err := saveInLocalDatabase(s, shortURL, originalURL); err != nil {
		return err
	}

	return nil
}

func saveInLocalDatabase(s *InMemory, shortURL, originalURL string) error {
	s.local.maxRecord = s.local.maxRecord + 1

	var tmpRecord tmpStorage
	tmpRecord.ID = strconv.Itoa(s.local.maxRecord)
	tmpRecord.ShortURL = shortURL
	tmpRecord.OriginalURL = originalURL

	data, err := json.Marshal(&tmpRecord)
	if err != nil {
		log.Fatalf("error: reading from json: %d", err)
	}

	data = append(data, '\n')
	_, err = s.local.file.Write(data)
	if err != nil {
		log.Fatalf("error: writing to json file: %d", err)
	}

	return nil
}

func (s *InMemory) GetURL(id string) (string, bool) {
	originalURL, ok := s.dataURL[id]
	return originalURL, ok
}
