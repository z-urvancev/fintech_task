package repository

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Repository interface {
	GetByShort(short string) (string, error)
	GetByURL(url string) (string, error)
	Insert(url, short string) error
}
