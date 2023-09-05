package service

type UserRepository interface {
	Create() error
}

type SlugRepository interface {
	Create() error
	Remove() error
	GetFromUser() error
	AddUser() error
}

type App interface {
	CreateSlug() error
	RemoveSlug() error
	GetSlugs() error
	AddUserToSlug() error

	CreateUser() error
}

type app struct {
	userrepo UserRepository
	slugrepo SlugRepository
}

func NewApp(slugRepo SlugRepository, userRepo UserRepository) App {
	return &app{
		slugrepo: slugRepo,
		userrepo: userRepo,
	}
}

func (a *app) CreateSlug() error    { return nil }
func (a *app) RemoveSlug() error    { return nil }
func (a *app) GetSlugs() error      { return nil }
func (a *app) AddUserToSlug() error { return nil }

func (a *app) CreateUser() error { return nil }
