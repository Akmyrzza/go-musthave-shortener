package pgsql

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"strings"
	"time"
)

type StoreDB struct {
	DB *sql.DB
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

func InitDatabase(DatabasePath string) (service.Repository, error) {
	if DatabasePath == "" {
		return nil, errors.New("error database path empty")
	}

	db, err := sql.Open("pgx", DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	d, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to return an iofs driver: %w", err)
	}

	//databaseURL := parseDatabasePath(DatabasePath)
	m, err := migrate.NewWithSourceInstance("iofs", d, DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get a new migrate instance: %w", err)
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, fmt.Errorf("failed to apply migrations to the DB: %w", err)
		}
	}

	storeDB := new(StoreDB)
	storeDB.DB = db

	return storeDB, nil
}

func parseDatabasePath(databasePath string) string {
	parts := strings.Split(databasePath, " ")
	params := make(map[string]string)

	for _, part := range parts {
		p := strings.SplitN(part, "=", 2)
		params[p[0]] = p[1]
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", params["user"], params["password"], params["host"], params["port"], params["dbname"], params["sslmode"])
	return databaseURL
}

func (s *StoreDB) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := s.DB.PingContext(ctx); err != nil {
		return fmt.Errorf("pinging db-pgsql: %w", err)
	}
	return nil
}

func (s *StoreDB) CreateShortURL(originalURL, shortURL string) (string, error) {
	query := `INSERT INTO urls (originalURL, shortURL) VALUES ($1, $2) ON CONFLICT (originalURL) DO UPDATE SET originalURL=$1 RETURNING shortURL`

	var id string
	err := s.DB.QueryRow(query, originalURL, shortURL).Scan(&id)
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

func (s *StoreDB) CreateShortURLs(urls []model.ReqURL) ([]model.ReqURL, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("transaction error: %w", err)
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("error rollback: %d", err)
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO urls (originalURL, shortURL)" + "VALUES($1, $2)")
	if err != nil {
		return nil, fmt.Errorf("tx query error: %w", err)
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("error statement: %d", err)
		}
	}()

	for i, v := range urls {
		_, err := stmt.Exec(v.OriginalURL, v.ShortURL)
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
