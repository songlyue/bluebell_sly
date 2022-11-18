package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

var mySecret = []byte("嗯嗯好大")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

const TokenExpireDuration = time.Hour * 24 * 365

func GenToken(userId uint64) (aToken, rToken string, err error) {
	claims := MyClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "songlongyue",
		},
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)

	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "songlongyue",
	}).SignedString(mySecret)

	return

}

func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("invalid token")
	}
	return
}

func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserId)
	}
	return

}
