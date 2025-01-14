package gomeas

import (
	"math/rand/v2"
)

const (
	CELSIUS    = `°C`
	FAHRENHEIT = `°F`
	HUMIDITY   = `%`
)

/* Message interface for interacting with the MQTT connection
 * Allows anyon to implemnt custom Message types if the provided
 * Messages are insufficient
 */
type Message interface {
	Marshal() ([]byte, error)
	GetTopic() ([]byte, error)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GenerateId creates a random alpha-numeric string of a given length n
func GenerateId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

func randInt16() uint16 {
	var a = rand.Uint32()
	a %= (65535 - 1)
	a += 1
	return uint16(a)
}
