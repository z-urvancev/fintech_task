package impl

import (
	"fintech/pkg/errors"
	"sync"
)

type InMemoryRepo struct {
	store map[string]string
	mutex *sync.RWMutex
}

func NewInMemoryRepo(store map[string]string) *InMemoryRepo {
	return &InMemoryRepo{store: store}
}

func (imr *InMemoryRepo) GetByShort(short string) (string, error) {
	imr.mutex.RLock()
	defer imr.mutex.RUnlock()
	url, ok := imr.store[short]
	if !ok {
		return "", errors.ErrURLNotFound
	}
	return url, nil
}

func (imr *InMemoryRepo) GetByURL(url string) (string, error) {
	imr.mutex.RLock()
	defer imr.mutex.RUnlock()
	for short, elem := range imr.store {
		if elem == url {
			return short, nil
		}
	}
	return "", errors.ErrURLNotFound
}

func (imr *InMemoryRepo) Insert(url, short string) error {
	imr.mutex.Lock()
	defer imr.mutex.Unlock()
	imr.store[short] = url
	return nil
}
