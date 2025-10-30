package crack

import (
    "crypto/sha512"
    "encoding/hex"
)

func SHA512Hash(text string) string {
    hasher := sha512.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
