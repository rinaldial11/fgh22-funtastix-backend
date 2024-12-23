package libs

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"github.com/joho/godotenv"
)

func GetMD5Hash() string {
	godotenv.Load()
	secret := os.Getenv("JWT_SECRET")
	hasher := md5.New()
	hasher.Write([]byte(secret))
	return hex.EncodeToString(hasher.Sum(nil))
}
