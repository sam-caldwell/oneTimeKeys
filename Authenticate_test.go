package oneTimeKeys

import (
	"testing"
)

func TestOTk_Authenticate(t *testing.T) {
	var key []OneTimeKey
	t.Run("setup the environment", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			key = append(key, GetNewOneTimeKey())
		}
	})
	t.Run("test", func(t *testing.T) {
		for i, _ := range key {
			if ok := Authenticate(key[i]); !ok {
				t.Fatalf("Failed to authenticate %d", i)
			}
			if ok := Authenticate(key[i]); ok {
				t.Fatalf("key %d should be deleted", i)
			}
		}
	})
}
