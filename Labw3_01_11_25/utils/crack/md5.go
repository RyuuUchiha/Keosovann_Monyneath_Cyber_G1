package crack

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
