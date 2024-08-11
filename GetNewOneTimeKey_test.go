package oneTimeKeys

import (
	"bytes"
	"fmt"
	"testing"
)

func TestOtk_GetNewOneTimeKey(t *testing.T) {
	t.Run("test size", func(t *testing.T) {
		otk := GetNewOneTimeKey()
		if len(otk) < 1024 {
			t.Fatalf("otk should be 1024 bytes")
		}
		otkQueue.Pop()
	})
	t.Run("test uniqueness", func(t *testing.T) {
		for pass := 0; pass < 100; pass++ {
			t.Run(fmt.Sprintf("pass:%d", pass), func(t *testing.T) {
				otk1 := GetNewOneTimeKey()
				otk2 := GetNewOneTimeKey()

				// Basic check to see if two consecutive keys are different
				if bytes.Equal(otk1[:], otk2[:]) {
					t.Fatalf("otk should be random; got two identical keys\n"+
						"otk1: %v\n"+
						"otk2: %v", otk1, otk2)
				}
				go otkQueue.Pop()
				otkQueue.Pop()
			})
		}
	})
}
