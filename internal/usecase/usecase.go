package usecase

type UseCase interface {
	GetURLByShort(short string) (string, error)
	GenerateShortURL(url string) (string, error)
}
