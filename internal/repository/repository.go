package repository

type Repository interface {
	GetByShort(short string) (string, error)
	GetByURL(url string) (string, error)
	Insert(url, short string) error
}
