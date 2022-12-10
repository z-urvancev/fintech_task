package abbreviator

import (
	"fintech/pkg/errors"
	"math/rand"
	"strings"
	"time"
)

var abbreviateMap = map[string]struct{}{}
var acceptableSymbols = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_"

func Generate() (string, error) {
	rand.Seed(time.Now().UnixNano())
	var (
		ok       = true
		result   = ""
		attempts = 0
	)
	for ok && attempts < 10 {
		result = generateShortUrl()
		_, ok = abbreviateMap[result]
		attempts++
	}
	if attempts == 10 {
		return "", errors.ErrCannotGenerateShort
	}
	return result, nil
}

func generateShortUrl() string {
	builder := strings.Builder{}
	builder.Grow(10)
	for i := 0; i < 10; i++ {
		num := rand.Int() % len(acceptableSymbols)
		builder.Write([]byte{acceptableSymbols[num]})
	}
	return builder.String()
}
