package oneTimeKeys

import "testing"

func TestOneTimeKeyType(t *testing.T) {
	o := OneTimeKey{}
	if len(o) != OneTimeKeySize {
		t.Fatal("OneTimeKey is supposed to be OneTimeKeySize in length")
	}
	if len(o) != 1024 {
		t.Fatal("OneTimeKey is supposed to be 1024 in length")
	}
}
