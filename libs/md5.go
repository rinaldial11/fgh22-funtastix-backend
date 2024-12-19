package libs

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(secret string) string {
	hasher := md5.New()
	hasher.Write([]byte(secret))
	return hex.EncodeToString(hasher.Sum(nil))
}
