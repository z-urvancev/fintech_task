package impl

import (
	"fintech/internal/repository"
	"fintech/internal/usecase/abbreviator"
	"fintech/pkg/errors"
)

type UseCaseImpl struct {
	repository repository.Repository
}

func NewUseCase(repository repository.Repository) *UseCaseImpl {
	return &UseCaseImpl{repository: repository}
}

func (u *UseCaseImpl) GetURLByShort(short string) (string, error) {
	url, getErr := u.repository.GetByShort(short)
	if getErr != nil {
		return "", getErr
	}

	return url, nil
}

func (u *UseCaseImpl) GenerateShortURL(url string) (string, error) {
	_, getErr := u.repository.GetByURL(url)
	if getErr == nil {
		return "", errors.ErrAlreadyAbbreviated
	}
	if getErr != errors.ErrURLNotFound {
		return "", getErr
	}

	short, abbreviateErr := abbreviator.Generate()
	if abbreviateErr != nil {
		return "", abbreviateErr
	}

	insertErr := u.repository.Insert(url, short)
	if insertErr != nil {
		return "", insertErr
	}

	return short, nil
}