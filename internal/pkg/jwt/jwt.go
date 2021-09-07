package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)
var (
	SecretKey = []byte("secret")
)

// ユーザのトークン生成
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// ユーザネームをclaims(請求)に保存
	claims["username"] = username
	// 有効期限を24時間に設定
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// トークンの受け取り
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// リクエストを送信したユーザの特定
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
