package repository

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

type LocalRepository struct {
	file      *os.File
	useFile   bool
	dataURL   map[string]string
	maxRecord int
}

type tmpStorage struct {
	ID          string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewLocalRepository(filePath string) (*LocalRepository, error) {
	if filePath == "" {
		return &LocalRepository{
			file:      nil,
			useFile:   false,
			dataURL:   make(map[string]string),
			maxRecord: 0,
		}, nil
	}

	fileDB, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	newDecoder := json.NewDecoder(fileDB)
	dataURL := make(map[string]string)
	maxRecord := 0

	for {
		var tmpRecord tmpStorage

		if err := newDecoder.Decode(&tmpRecord); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		dataURL[tmpRecord.ShortURL] = tmpRecord.OriginalURL
		maxRecord, err = strconv.Atoi(tmpRecord.ID)
		if err != nil {
			return nil, err
		}
	}

	return &LocalRepository{
		file:      fileDB,
		useFile:   true,
		dataURL:   dataURL,
		maxRecord: maxRecord,
	}, nil
}

func (s *LocalRepository) CreateID(shortURL, originalURL string) error {
	_, found := s.dataURL[shortURL]
	if found {
		return errors.New("this id already exist")
	}

	s.dataURL[shortURL] = originalURL
	s.maxRecord = s.maxRecord + 1

	if s.useFile == false {
		return nil
	}

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
	originalURL, ok := s.dataURL[id]
	return originalURL, ok
}
