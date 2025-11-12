package movie

import "fmt"

type Movie struct {
	ID    int
	Title string
	Year  int
}

type Store interface {
	GetMovieByID(id int) (Movie, error)
}

type Service struct {
	store Store
}

func NewService(s Store) *Service {
	return &Service{store: s}
}

func (svc *Service) GetMovieInfo(id int) (string, error) {
	m, err := svc.store.GetMovieByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to get movie: %w", err)
	}
	if id < 1980 {
		return `Classic`, nil
	}

	return fmt.Sprintf("%s (%d)", m.Title, m.Year), nil
}
