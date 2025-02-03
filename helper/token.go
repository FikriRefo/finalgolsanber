package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("mysecretkey")

type JWTClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int) (string, error) {
	claims := JWTClaims{
		ID: id, // Menggunakan parameter id, bukan user.ID yang tidak dideklarasikan
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, errors.New("invalid token signature") // Perbaikan: return 0 bukan nil
		}
		return 0, errors.New("your token is expired") // Perbaikan: return 0 bukan nil
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return 0, errors.New("your token is expired") // Perbaikan: return 0 bukan nil
	}

	return claims.ID, nil
}
