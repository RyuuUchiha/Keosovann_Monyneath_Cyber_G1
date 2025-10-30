package crack

import (
    "crypto/sha1"
    "encoding/hex"
)

// SHA1Hash takes a plaintext string and returns its SHA1 hash
func SHA1Hash(text string) string {
    hasher := sha1.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
