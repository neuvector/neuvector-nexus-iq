// +build e2e

package e2e

import (
	"github.com/cenkalti/backoff/v4"
	"math/rand"
	"time"
)

// Initialize random number source
var rng = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

var hexLetters = []rune("0123456789abcdef")

func randString(n int, letters []rune) string {
	b := make([]rune, n)
	l := len(letters)
	for i := range b {
		b[i] = letters[rng.Intn(l)]
	}
	return string(b)
}

func randHex(n int) string {
	return randString(n, hexLetters)
}

func newTestId() string {
	//return uuid.New().String()
	return randHex(16)
}

func retryWithConstantBackoff(operation backoff.Operation, backoffInterval time.Duration, maxRetries uint64) error {
	return backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewConstantBackOff(backoffInterval), maxRetries))
}
