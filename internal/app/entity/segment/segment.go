package segment

type Segment struct {
	ID   int64  `db:"id"`
	Slug string `db:"segment"`
}
