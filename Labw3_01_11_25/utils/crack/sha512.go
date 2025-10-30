package crack

import (
    "crypto/sha512"
    "encoding/hex"
)

// SHA512Hash takes a plaintext string and returns its SHA512 hash
func SHA512Hash(text string) string {
    hasher := sha512.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
