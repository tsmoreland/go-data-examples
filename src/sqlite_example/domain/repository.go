package domain

type Repository interface {
	Migrate() error
	Close() error
}
