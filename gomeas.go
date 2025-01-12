package gomeas

import (
	"math/rand/v2"
)

const (
	CELSIUS    = `°C`
	FAHRENHEIT = `°F`
)

func randInt16() uint16 {
	var a = rand.Uint32()
	a %= (65535 - 1)
	a += 1
	return uint16(a)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}
