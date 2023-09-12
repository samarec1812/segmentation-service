package service

type UserRepository interface {
	Create() error
}

type SegmentRepository interface {
	Create() error
	Remove() error
	GetFromUser() error
	AddUser() error
}

type App interface {
	CreateSegment(string) error
	RemoveSegment() error
	GetSegments() error
	AddUserToSegment() error

	CreateUser() error
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

func (a *app) CreateSegment(slug string) error {
	return a.sgrepo.Create()
}

func (a *app) RemoveSegment() error    { return nil }
func (a *app) GetSegments() error      { return nil }
func (a *app) AddUserToSegment() error { return nil }

func (a *app) CreateUser() error { return nil }
