package services

import (
	"employee_manage/config"
	"employee_manage/constant"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenStruct struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}
type TokenPayload struct {
	UserID int `json:"user_id"`
}

type MyCustomClaims struct {
	Payload TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func SignToken(tokenType string, payload TokenPayload) (tokenStruct TokenStruct, err error) {

	var expired int
	var secret string

	if tokenType == constant.AccessToken {
		expired = config.ConfigApp.JwtConfig.ExpiredAccessToken
		secret = config.ConfigApp.JwtConfig.SecretAccessToken
	} else {
		expired = config.ConfigApp.JwtConfig.ExpiredRefreshToken
		secret = config.ConfigApp.JwtConfig.SecretRefreshToken
	}

	claims := MyCustomClaims{
		payload,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expired) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return
	}

	tokenStruct = TokenStruct{
		Type:  tokenType,
		Token: t,
	}

	return
}
