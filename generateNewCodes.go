package oneTimeKeys

import (
	"crypto/rand"
	"crypto/sha256"
	"time"
)

// generateNewCode - Generate a random one-time key for authenticating discovery
func generateNewCode() OneTimeKey {
	const sz = OneTimeKeySize - 32
	var errorCount int
	randomBytes := make([]byte, sz)
	if _, err := rand.Read(randomBytes); err != nil {
		errorCount++
		if errorCount > MaxAuthCodeGeneratorErrors {
			panic(err)
		}
		time.Sleep(10 * time.Second)
	} else {
		errorCount = 0
	}
	hash := sha256.Sum256(randomBytes[:])
	return OneTimeKey(append(randomBytes[:], hash[:]...))
}
