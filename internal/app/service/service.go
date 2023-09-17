package service

import (
	"context"
	"errors"

	"github.com/samarec1812/segmentation-service/internal/app/entity/segment"
)

type UserRepository interface {
	Create(context.Context) error
}

type SegmentRepository interface {
	Create(context.Context, segment.Segment) error
	Remove(context.Context, string) (int64, error)
	GetFromUser(context.Context, int64) ([]segment.Segment, error)
	AddUser(context.Context, int64, []string, []string) error
}

type App interface {
	CreateSegment(context.Context, string) error
	RemoveSegment(context.Context, string) error
	GetSegments(context.Context, int64) ([]string, error)
	AddUserToSegment(context.Context, int64, []string, []string) error

	CreateUser(context.Context) error
}

type app struct {
	userrepo UserRepository
	sgrepo   SegmentRepository
}

func NewApp(sgRepo SegmentRepository, userRepo UserRepository) App {
	return &app{
		sgrepo:   sgRepo,
		userrepo: userRepo,
	}
}

func (a *app) CreateSegment(ctx context.Context, slug string) error {
	sg := segment.Segment{
		Slug: slug,
	}
	return a.sgrepo.Create(ctx, sg)
}

func (a *app) RemoveSegment(ctx context.Context, slug string) error {
	res, err := a.sgrepo.Remove(ctx, slug)
	if err != nil {
		return err
	}

	if res < 1 {
		return errors.New("no deleted row")
	}

	return nil
}

func (a *app) GetSegments(ctx context.Context, userID int64) ([]string, error) {
	res, err := a.sgrepo.GetFromUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var segments []string
	for _, seg := range res {
		segments = append(segments, seg.Slug)
	}

	return segments, nil
}

func (a *app) AddUserToSegment(ctx context.Context, userID int64, addSegments, removeSegments []string) error {
	err := a.sgrepo.AddUser(ctx, userID, addSegments, removeSegments)
	return err
}

func (a *app) CreateUser(ctx context.Context) error { return nil }
