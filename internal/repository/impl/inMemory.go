package impl

import "fintech/pkg/errors"

type InMemoryRepo struct {
	store map[string]string
}

func NewInMemoryRepo(store map[string]string) *InMemoryRepo {
	return &InMemoryRepo{store: store}
}

func (imr *InMemoryRepo) GetByShort(short string) (string, error) {
	url, ok := imr.store[short]
	if !ok {
		return "", errors.ErrURLNotFound
	}
	return url, nil
}

func (imr *InMemoryRepo) GetByURL(url string) (string, error) {
	for short, elem := range imr.store {
		if elem == url {
			return short, nil
		}
	}
	return "", errors.ErrURLNotFound
}

func (imr *InMemoryRepo) Insert(url, short string) error {
	imr.store[short] = url
	return nil
}
