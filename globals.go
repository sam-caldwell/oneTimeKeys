package oneTimeKeys

import (
	"bytes"
	"github.com/sam-caldwell/genericQueue"
)

// otkQueue - Queue of known one-time keys
var otkQueue = *genericQueue.NewGenericQueue[OneTimeKey](func(a, b OneTimeKey) int {
	//compareFunction for the queue.
	return bytes.Compare(a[:], b[:])
})
