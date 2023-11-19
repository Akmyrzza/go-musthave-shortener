package repository

type Repository interface {
	CreateID(id, originalURL string)
	GetURL(id string) (string, bool)
}
