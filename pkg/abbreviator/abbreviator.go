package abbreviator

import (
	"fintech/pkg/errors"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const acceptableSymbols = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_"

type AbbreviateGenerator struct {
	store map[string]struct{}
	mutex *sync.RWMutex
}

func NewAbbreviateGenerator() *AbbreviateGenerator {
	rand.Seed(time.Now().UnixNano())
	return &AbbreviateGenerator{store: make(map[string]struct{}), mutex: new(sync.RWMutex)}
}

func (ag *AbbreviateGenerator) Generate() (string, error) {
	ag.mutex.Lock()
	defer ag.mutex.Unlock()
	var (
		ok       = true
		result   = ""
		attempts = 0
	)
	for ok && attempts < 10 {
		result = ag.generateShortUrl()
		_, ok = ag.store[result]
		attempts++
	}
	if attempts == 10 {
		return "", errors.ErrCannotGenerateShort
	}
	ag.store[result] = struct{}{}
	return result, nil
}

func (ag *AbbreviateGenerator) generateShortUrl() string {
	builder := strings.Builder{}
	builder.Grow(10)
	for i := 0; i < 10; i++ {
		num := rand.Int() % len(acceptableSymbols)
		builder.Write([]byte{acceptableSymbols[num]})
	}
	return builder.String()
}
