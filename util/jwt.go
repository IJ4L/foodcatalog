package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	Id      int       `json:"id"`
	Expires time.Time `json:"expires"`
}

var secretKey []byte
var tokenExpiry time.Duration

func InitToken(secret string, expiredToken int) {
	secretKey = []byte(secret)
	tokenExpiry = time.Duration(expiredToken) * time.Hour
}

func NewJWT(id int) JWT {
	return JWT{
		Id:      id,
		Expires: time.Now().Add(tokenExpiry),
	}
}

func (j JWT) GenerateToken() (string, error) {
	expires := time.Now().Add(tokenExpiry)

	claims := jwt.MapClaims{
		"id":      j.Id,
		"expires": expires.Format(time.RFC3339),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokString, err := token.SignedString(secretKey)
	return tokString, err
}

func VerifyToken(tokString string) (JWT, error) {

	jwtToken, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return JWT{}, fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return JWT{}, fmt.Errorf("invalid token")
	}

	id, ok := claims["id"].(float64)
	if !ok {
		return JWT{}, fmt.Errorf("invalid token ID")
	}

	expiresStr, ok := claims["expires"].(string)
	if !ok {
		return JWT{}, fmt.Errorf("missing expiry time in token")
	}

	expiresTime, err := time.Parse(time.RFC3339, expiresStr)
	if err != nil {
		return JWT{}, fmt.Errorf("failed to parse expiry time: %v", err)
	}

	if time.Now().After(expiresTime) {
		return JWT{}, fmt.Errorf("token expired")
	}

	return JWT{Id: int(id), Expires: expiresTime}, nil
}
