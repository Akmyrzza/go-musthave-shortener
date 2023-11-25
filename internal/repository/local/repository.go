package local

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type LocalRepository struct {
	file      *os.File
	dataURL   map[string]string
	maxRecord int
}

type TmpStorage struct {
	ID          string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewLocalRepository(filePath string) (*LocalRepository, error) {
	if filePath == "" {
		return &LocalRepository{
			file:      nil,
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
		var tmpRecord TmpStorage

		if err := newDecoder.Decode(&tmpRecord); err != nil {
			break
		}

		dataURL[tmpRecord.ShortURL] = tmpRecord.OriginalURL
		maxRecord, err = strconv.Atoi(tmpRecord.ID)
		if err != nil {
			return nil, err
		}
	}

	return &LocalRepository{
		file:      fileDB,
		dataURL:   dataURL,
		maxRecord: maxRecord,
	}, nil
}

func (s *LocalRepository) CreateID(shortURL, originalURL string) {
	s.dataURL[shortURL] = originalURL
	s.maxRecord = s.maxRecord + 1

	if s.file == nil {
		return
	}

	var tmpRecord TmpStorage
	tmpRecord.ID = strconv.Itoa(s.maxRecord)
	tmpRecord.ShortURL = shortURL
	tmpRecord.OriginalURL = originalURL

	data, err := json.Marshal(&tmpRecord)
	if err != nil {
		log.Fatalf("error: reading from json: %w", err)
		return
	}

	data = append(data, '\n')
	_, err = s.file.Write(data)
	if err != nil {
		log.Fatalf("error: writing to json file: %w", err)
		return
	}
}

func (s *LocalRepository) GetURL(id string) (string, bool) {
	originalURL, ok := s.dataURL[id]
	return originalURL, ok
}
