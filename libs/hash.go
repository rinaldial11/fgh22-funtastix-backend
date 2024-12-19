package libs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pilinux/argon2"
)

func CreateHash(password string) string {
	godotenv.Load()
	var SECRET_KEY = os.Getenv("SECRET_KEY")
	hasher, _ := argon2.CreateHash(password, string(SECRET_KEY), argon2.DefaultParams)
	return hasher
}

func HashValidator(password string, hash string) bool {
	godotenv.Load()
	var SECRET_KEY = os.Getenv("SECRET_KEY")
	isValid, _ := argon2.ComparePasswordAndHash(password, SECRET_KEY, hash)
	return isValid
}
