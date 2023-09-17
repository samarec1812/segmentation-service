package segment

type Segment struct {
	ID   int64  `db:"id"`
	Slug string `db:"segment"`
}

func (s *Segment) GetSegmentDBRecord() map[string]any {
	return map[string]any{
		"slug": s.Slug,
	}
}
