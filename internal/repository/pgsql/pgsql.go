package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

type StoreDB struct {
	DB *sql.DB
}

func InitDatabase(DatabasePath string) (service.Repository, error) {
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
	err := tableExist(db, tableName)
	if err != nil {
		err := createTable(db, tableName)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func tableExist(db *sql.DB, tableName string) error {
	query := `SELECT COUNT(*) from urls WHERE table_name = $1`
	err := db.QueryRow(query, tableName)
	if err != nil {
		return fmt.Errorf("table does not exist: %w", err)
	}

	return nil
}

func createTable(db *sql.DB, tableName string) error {
	query := `CREATE TABLE ` + tableName + ` (
				id SERIAL PRIMARY KEY,
				originalURL VARCHAR(255) NOT NULL,
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

func (s *StoreDB) CreateShortURL(originalURL, shortURL string) error {
	query := `INSERT INTO urls (originalURL, shortURL) VALUES ($1, $2)`

	_, err := s.DB.Exec(query, originalURL, shortURL)
	if err != nil {
		return fmt.Errorf("error: db query exec: %w", err.Error())
	}

	return nil
}

func (s *StoreDB) GetOriginalURL(shortURL string) (string, error) {
	var url string

	query := `SELECT originalURL from urls WHERE shortURL = $1`

	row := s.DB.QueryRow(query, shortURL)

	err := row.Scan(&url)
	if err != nil {
		return "", fmt.Errorf("error: db query: %w", err.Error())
	}

	return url, nil
}
