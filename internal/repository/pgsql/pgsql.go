package pgsql

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type StoreDB struct {
	DB *pgxpool.Pool
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

func InitDatabase(DatabasePath string) (service.Repository, error) {
	if DatabasePath == "" {
		return nil, errors.New("error database path empty")
	}

	//db, err := sql.Open("pgx", DatabasePath)
	db, err := pgxpool.New(context.Background(), DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	d, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to return an iofs driver: %w", err)
	}

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

func (s *StoreDB) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := s.DB.Ping(ctx); err != nil {
		return fmt.Errorf("pinging db-pgsql: %w", err)
	}
	return nil
}

func (s *StoreDB) CreateShortURL(ctx context.Context, originalURL, shortURL string) (string, error) {
	var id string

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("transaction error: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				log.Printf("error rollback: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(ctx); cmErr != nil {
				log.Printf("error commit: %v", cmErr)
			}
		}

	}()

	queryGet := `SELECT shortURL FROM urls WHERE originalURL = $1`
	result := tx.QueryRow(ctx, queryGet, originalURL)
	if err := result.Scan(&id); err != nil {
		if err == pgx.ErrNoRows {
			query := `INSERT INTO urls (originalURL, shortURL) VALUES ($1, $2)`
			_, err := tx.Exec(ctx, query, originalURL, shortURL)
			if err != nil {
				return "", fmt.Errorf("error: db query exec: %w", err)
			}
			return shortURL, nil
		}
		return "", fmt.Errorf("db query error: %w", err)
	}
	return id, cerror.ErrAlreadyExist
}

func (s *StoreDB) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	var url string

	query := `SELECT originalURL from urls WHERE shortURL = $1`

	row := s.DB.QueryRow(ctx, query, shortURL)

	err := row.Scan(&url)
	if err != nil {
		return "", fmt.Errorf("error: db query: %w", err)
	}

	return url, nil
}

func (s *StoreDB) CreateShortURLs(ctx context.Context, urls []model.ReqURL) ([]model.ReqURL, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("transaction error: %w", err)
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				log.Printf("error rollback: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(ctx); cmErr != nil {
				log.Printf("error commit: %v", cmErr)
			}
		}

	}()

	batch := &pgx.Batch{}

	for i, v := range urls {
		batch.Queue("INSERT INTO urls (originalURL, shortURL) VALUES($1, $2)", v.OriginalURL, v.ShortURL)
		urls[i].OriginalURL = ""
	}

	br := tx.SendBatch(ctx, batch)
	defer func() {
		if err := br.Close(); err != nil {
			log.Printf("error batch close: %d", err)
		}
	}()

	_, err = br.Exec()
	if err != nil {
		return nil, fmt.Errorf("batch execution error: %w", err)
	}

	return urls, nil
}
