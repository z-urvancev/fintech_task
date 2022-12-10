package abbreviator

import (
	"math/rand"
	"strings"
	"time"
)

var abbreviateMap = map[string]struct{}{}
var acceptableSymbols = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_"

func Generate() (string, error) {
	rand.Seed(time.Now().UnixNano())
	ok := true
	for ok {
		_, ok = abbreviateMap[generateShortUrl()]
	}
	return "", nil
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
