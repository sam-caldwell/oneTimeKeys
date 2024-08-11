package oneTimeKeys

import (
	"crypto/rand"
)

// GetNewOneTimeKey - Create, Register and return a new one-time key
func GetNewOneTimeKey() OneTimeKey {
	randomBytes := make([]byte, OneTimeKeySize)
	_, _ = rand.Read(randomBytes)
	otk := OneTimeKey(randomBytes[0:OneTimeKeySize])
	otkQueue.Push(otk)
	return otk
}
