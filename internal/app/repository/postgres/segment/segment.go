package segment

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/samarec1812/segmentation-service/internal/app/entity/segment"
)

const (
	segmentsTable     = "segments"
	userSegmentsTable = "user_segments"
)

type Repository struct {
	db *sql.DB
}

func NewSegmentRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, sg segment.Segment) error {
	query, args, err := sq.Insert(segmentsTable).SetMap(sg.GetSegmentDBRecord()).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Remove(ctx context.Context, slug string) (int64, error) {
	query, args, err := sq.Delete(segmentsTable).Where(sq.Eq{"slug": slug}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, err
	}
	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}

func (r *Repository) GetFromUser(ctx context.Context, userID int64) ([]segment.Segment, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := builder.Select("slug").
		From(segmentsTable + " s").
		Join(fmt.Sprintf("%s us ON us.segment_id = s.id", userSegmentsTable)).
		Where(sq.Eq{"us.user_id": userID}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("error search segments: %w", err)
	}

	segments := make([]segment.Segment, 0)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error search segments: %w", err)
	}

	for rows.Next() {
		var seg segment.Segment
		err = rows.Scan(&seg.Slug)
		if err != nil {
			return nil, fmt.Errorf("error search segments: %w", err)
		}

		segments = append(segments, seg)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error search segments: %w", err)
	}

	return segments, nil
}

func (r *Repository) AddUser(ctx context.Context, userID int64, addSegments, _ []string) error {

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qs := builder.Insert(userSegmentsTable).Columns("user_id", "segment_id")
	subquery := fmt.Sprintf("(select id from %s where slug=?)", segmentsTable)

	values := []interface{}{}

	for _, slug := range addSegments {
		values = append(values, userID, sq.Expr(subquery, slug))
		qs = qs.Values(values...)
		values = nil
	}

	query, args, err := qs.ToSql()
	if err != nil {
		return nil
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
