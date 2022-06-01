package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

var jwtKey = []byte("systemMoniter-Server")

type JWTClaim struct {
	Host string `json:"host"`
	jwt.StandardClaims
}

func GenerateJWT(host string, username string) (tokenString string, err error) {
	zap.L().Info("Create JWT Token: ")
	expirationTime := time.Now().Add(3 * time.Hour * time.Duration(1))
	claims := &JWTClaim{
		Host: host,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), // 签发时间
			NotBefore: time.Now().Unix(), // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	zap.L().Info("Create JWT Finish: " + tokenString)
	return
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
	zap.L().Info("Validate JWT Token: ")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		zap.L().Info("Validate JWT Token Err: ", zap.Error(err))
		return nil, err
	}
	if token != nil {
		claims, ok := token.Claims.(*JWTClaim)
		if !ok {
			zap.L().Info("Validate JWT Token Err: couldn't parse claims")
			err = errors.New("couldn't parse claims")
			return nil, err
		}

		if claims.ExpiresAt < time.Now().Local().Unix() {
			zap.L().Info("Validate JWT Token Err: token expired")
			err = errors.New("token expired")
			return nil, err
		}

		if token.Valid && ok {
			return claims, nil
		}
	} else {
		zap.L().Info("Validate JWT Token Err: token is null")
	}

	zap.L().Info("Validate JWT Token Err: unknow error")
	return nil, err

}
