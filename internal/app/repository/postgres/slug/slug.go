package slug

import "database/sql"

type SlugRepository struct {
	db *sql.DB
}

func NewSlugRepository(db *sql.DB) *SlugRepository {
	return &SlugRepository{db: db}
}

func (s *SlugRepository) Create() error      { return nil }
func (s *SlugRepository) Remove() error      { return nil }
func (s *SlugRepository) GetFromUser() error { return nil }
func (s *SlugRepository) AddUser() error     { return nil }
