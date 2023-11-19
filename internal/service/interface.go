package service

type Service interface {
	CreateID(originalURL string) string
	GetURL(id string) (string, bool)
}
