package crypto

import(
	"time"
	"crypto/sha256"
	"encoding/hex"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/model"
)

var KEY string

func init() {
	KEY = time.Now().String()
}

func CreateHashWithPass(password string) string {
	byteHash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(byteHash[:])
}

func CreateToken(user model.DBUser) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
    claims["id"] = user.ID
    claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	
	return token.SignedString([]byte(KEY))
}
