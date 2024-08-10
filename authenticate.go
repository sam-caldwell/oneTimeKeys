package oneTimeKeys

import "bytes"

// Authenticate - verify a given claimed key exists in the GenericQueue and delete it if so before returning true.
// otherwise return false and do nothing else.
func Authenticate(claimedKey OneTimeKey) (ok bool) {
	ok = otkQueue.DeleteIfExists(claimedKey, bytes.Compare)
	return ok
}
