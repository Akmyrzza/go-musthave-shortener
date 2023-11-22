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

func NewLocalRepository(filePath string) *LocalRepository {
	if filePath == "" {
		return &LocalRepository{
			file:      nil,
			dataURL:   make(map[string]string),
			maxRecord: 0,
		}
	}
	fileDB, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error: opening db file: %d", err)
		return nil
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
		maxRecord, _ = strconv.Atoi(tmpRecord.ID)
	}

	return &LocalRepository{
		file:      fileDB,
		dataURL:   dataURL,
		maxRecord: maxRecord,
	}
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
		log.Fatalf("error: reading from json: %d", err)
		return
	}

	data = append(data, '\n')
	_, err = s.file.Write(data)
}

func (s *LocalRepository) GetURL(id string) (string, bool) {
	originalURL, ok := s.dataURL[id]
	return originalURL, ok
}
