package abbreviator

import (
	"math/rand"
	"strings"
	"time"
)

var abbreviateMap = map[string]struct{}{}
var acceptableSymbols = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_"

func Generate() string {
	rand.Seed(time.Now().UnixNano())
	var (
		ok     = true
		result = ""
	)
	for ok {
		result = generateShortUrl()
		_, ok = abbreviateMap[result]
	}
	return result
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
