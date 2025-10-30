package crack

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash returns the lowercase hex string of the MD5 sum of input.
func MD5Hash(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
