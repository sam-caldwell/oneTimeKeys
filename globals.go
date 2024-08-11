package oneTimeKeys

import (
	"bytes"
	"github.com/sam-caldwell/expiringQueue"
	"time"
)

// otkQueue - Queue of known one-time keys
var otkQueue = *expiringQueue.NewQueue[OneTimeKey](120*time.Second, func(a, b OneTimeKey) bool {
	//compareFunction for the queue.
	return bytes.Equal(a[:], b[:])
})
