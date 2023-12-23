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
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

type StoreDB struct {
	DB *pgx.Conn
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

func InitDatabase(DatabasePath string) (service.Repository, error) {
	if DatabasePath == "" {
		return nil, errors.New("error database path empty")
	}

	//db, err := sql.Open("pgx", DatabasePath)
	db, err := pgx.Connect(context.Background(), DatabasePath)
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

func (s *StoreDB) CreateShortURL(originalURL, shortURL string) (string, error) {
	var id string

	queryGet := `SELECT shortURL FROM urls WHERE originalURL = $1`
	result := s.DB.QueryRow(context.Background(), queryGet, originalURL)
	if err := result.Scan(&id); err != nil {
		if err == pgx.ErrNoRows {
			query := `INSERT INTO urls (originalURL, shortURL) VALUES ($1, $2)`
			_, err := s.DB.Exec(context.Background(), query, originalURL, shortURL)
			if err != nil {
				return "", fmt.Errorf("error: db query exec: %w", err)
			}
			return shortURL, nil
		}
		return "", fmt.Errorf("db query error: %w", err)
	}
	return id, cerror.ErrAlreadyExist
}

func (s *StoreDB) GetOriginalURL(shortURL string) (string, error) {
	var url string

	query := `SELECT originalURL from urls WHERE shortURL = $1`

	row := s.DB.QueryRow(context.Background(), query, shortURL)

	err := row.Scan(&url)
	if err != nil {
		return "", fmt.Errorf("error: db query: %w", err)
	}

	return url, nil
}

func (s *StoreDB) CreateShortURLs(urls []model.ReqURL) ([]model.ReqURL, error) {
	tx, err := s.DB.Begin(context.Background())
	if err != nil {
		return nil, fmt.Errorf("transaction error: %w", err)
	}
	defer func() {
		if err := tx.Rollback(context.Background()); err != nil {
			log.Printf("error rollback: %d", err)
		}
	}()

	batch := pgx.Batch{}

	for i, v := range urls {
		batch.Queue("INSERT INTO urls (originalURL, shortURL)"+"VALUES($1, $2)", v.OriginalURL, v.ShortURL)
		urls[i].OriginalURL = ""
	}

	br := tx.SendBatch(context.Background(), &batch)
	defer func() {
		if err := br.Close(); err != nil {
			log.Printf("error batch close: %d", err)
		}
	}()

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, fmt.Errorf("commit error: %w", err)
	}

	return urls, nil

	//stmt, err := tx.Prepare("INSERT INTO urls (originalURL, shortURL)" + "VALUES($1, $2)")
	//if err != nil {
	//	return nil, fmt.Errorf("tx query error: %w", err)
	//}
	//
	//defer func() {
	//	if err := stmt.Close(); err != nil {
	//		log.Printf("error statement: %d", err)
	//	}
	//}()
	//
	//for i, v := range urls {
	//	_, err := stmt.Exec(v.OriginalURL, v.ShortURL)
	//	if err != nil {
	//		return nil, fmt.Errorf("statement exec error: %w", err)
	//	}
	//	urls[i].OriginalURL = ""
	//}
	//
	//err = tx.Commit()
	//if err != nil {
	//	return nil, fmt.Errorf("commit error: %w", err)
	//}
	//
	//return urls, nil
}
