package usecase

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCase interface {
	GetURLByShort(short string) (string, error)
	GenerateShortURL(url string) (string, error)
}
