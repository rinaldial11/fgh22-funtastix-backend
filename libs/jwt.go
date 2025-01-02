package libs

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

type Claims struct {
	jwt.Claims
}

type ClaimsWithPayload struct {
	Claims
	UserID int
}

func GenerateToken(payload any) string {
	JWT_SECRET := []byte(GetMD5Hash())
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: JWT_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		fmt.Println(err)
	}
	baseInfo := jwt.Claims{
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}

	token, _ := jwt.Signed(sig).Claims(baseInfo).Claims(payload).Serialize()

	return token

}

func ValidateToken(head string) (map[string]interface{}, error) {
	JWT_SECRET := []byte(GetMD5Hash())
	token := strings.Split(head, " ")[1:][0]
	tok, _ := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
	// out := jwt.Claims{}
	out := make(map[string]interface{})
	err := tok.Claims(JWT_SECRET, &out)
	return out, err
}
