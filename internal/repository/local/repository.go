package local

type LocalRepository struct {
	dataURL map[string]string
}

func NewLocalRepository() *LocalRepository {
	return &LocalRepository{
		dataURL: make(map[string]string),
	}
}

func (s *LocalRepository) CreateID(id, originalURL string) {
	s.dataURL[id] = originalURL
}

func (s *LocalRepository) GetURL(id string) (string, bool) {
	originalURL, ok := s.dataURL[id]
	return originalURL, ok
}
