package oneTimeKeys

import (
	"bytes"
	"testing"
)

func TestOneTimeKeys_Global_variables(t *testing.T) {

	t.Run("basic test", func(t *testing.T) {
		if count := otkQueue.Count(); count != 0 {
			t.Fatalf("otkQueue should have no oneTimeKeys.  Got: %d", count)
		}
	})

	t.Run("make sure otk are unique", func(t *testing.T) {
		for i := 0; i < 1048576; i++ {
			var prev OneTimeKey
			var thisOne OneTimeKey
			var err error
			if thisOne = GetNewOneTimeKey(); err != nil {
				t.Fatalf("otkQueue should pop OTK (unique test)  Got: %s", err)
			}
			if ok := otkQueue.DeleteIfExists(thisOne); !ok {
				t.Fatalf("otkQueue should delete OTK.  Got: %d", otkQueue.Count())
			}
			if bytes.Equal(prev[:], thisOne[:]) {
				t.Fatalf("otkQueue should have oneTimeKeys (no repeats).\n"+
					"thisOne: %02x\n"+
					"   prev: %02x", thisOne, prev)
			}
			prev = thisOne
		}
	})

}
