package segment

import "database/sql"

type SegmentRepository struct {
	db *sql.DB
}

func NewSlugRepository(db *sql.DB) *SegmentRepository {
	return &SegmentRepository{db: db}
}

func (s *SegmentRepository) Create() error      { return nil }
func (s *SegmentRepository) Remove() error      { return nil }
func (s *SegmentRepository) GetFromUser() error { return nil }
func (s *SegmentRepository) AddUser() error     { return nil }
