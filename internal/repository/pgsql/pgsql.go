package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"time"
)

type StoreDB struct {
	DB *sql.DB
}

func InitDatabase(DatabasePath string) (service.Repository, error) {
	if DatabasePath == "" {
		return nil, errors.New("error database path empty")
	}

	db, err := sql.Open("pgx", DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	if err := initTables(db, "urls"); err != nil {
		return nil, err
	}

	storeDB := new(StoreDB)
	storeDB.DB = db

	return storeDB, nil
}

func initTables(db *sql.DB, tableName string) error {
	err := tableExist(db)
	if err != nil {
		errCreating := createTable(db, tableName)
		if errCreating != nil {
			return errCreating
		}

		return nil
	}

	return nil
}

func tableExist(db *sql.DB) error {
	var count int
	query := `SELECT COUNT(*) from urls`
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return fmt.Errorf("table does not exist: %w", err)
	}

	return nil
}

func createTable(db *sql.DB, tableName string) error {
	query := `CREATE TABLE ` + tableName + ` (
				id SERIAL PRIMARY KEY,
				userID VARCHAR(255) NOT NULL,
				originalURL VARCHAR(255) UNIQUE NOT NULL,
				shortURL VARCHAR(255) UNIQUE NOT NULL
				);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("table creating error: %w", err)
	}

	return nil
}

func (s *StoreDB) PingStore() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := s.DB.PingContext(ctx); err != nil {
		return fmt.Errorf("pinging db-pgsql: %w", err)
	}
	return nil
}

func (s *StoreDB) CreateShortURL(userID, originalURL, shortURL string) (string, error) {
	query := `INSERT INTO urls (userID, originalURL, shortURL) VALUES ($1, $2, $3) ON CONFLICT (originalURL) DO UPDATE SET originalURL=$2 RETURNING shortURL`

	var id string
	err := s.DB.QueryRow(query, userID, originalURL, shortURL).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error: db query exec: %w", err)
	}

	if id == shortURL {
		return "", nil
	}
	return id, nil
}

func (s *StoreDB) GetOriginalURL(shortURL string) (string, error) {
	var url string

	query := `SELECT originalURL from urls WHERE shortURL = $1`

	row := s.DB.QueryRow(query, shortURL)

	err := row.Scan(&url)
	if err != nil {
		return "", fmt.Errorf("error: db query: %w", err)
	}

	return url, nil
}

func (s *StoreDB) CreateShortURLs(userID string, urls []model.ReqURL) ([]model.ReqURL, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("transaction error: %w", err)
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("error rollback: %d", err)
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO urls (userID, originalURL, shortURL)" + "VALUES($1, $2, $3)")
	if err != nil {
		return nil, fmt.Errorf("tx query error: %w", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("error statement: %d", err)
		}
	}()

	for i, v := range urls {
		_, err := stmt.Exec(userID, v.OriginalURL, v.ShortURL)
		if err != nil {
			return nil, fmt.Errorf("statement exec error: %w", err)
		}
		urls[i].OriginalURL = ""
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("commit error: %w", err)
	}

	return urls, nil
}

func (s *StoreDB) GetAllURLs(userID string) ([]model.ResURL, error) {
	var data []model.ResURL

	query := `SELECT shorturl, originalurl FROM urls WHERE userID = $1`

	rows, err := s.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("database query exec error: %w", err)
	}

	for rows.Next() {
		var row model.ResURL

		err := rows.Scan(&row.ShortURL, &row.OriginalURL)
		if err != nil {
			return nil, fmt.Errorf("rows scanning error: %w", err)
		}

		data = append(data, row)
	}

	return data, nil
}
