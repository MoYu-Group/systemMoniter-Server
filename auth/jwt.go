package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

var jwtKey = []byte("systemMoniter-Server")

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (tokenString string, err error) {
	zap.L().Info("Start create token: ")
	expirationTime := time.Now().Add(3 * time.Hour * time.Duration(1))
	claims := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), // 签发时间
			NotBefore: time.Now().Unix(), // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		msg := "create token error: " + err.Error()
		zap.L().Error(msg)
	}
	zap.L().Info("create token success: " + tokenString)
	return
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
	zap.L().Info("Start Validating token : ")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		zap.L().Error("Validate token error: ", zap.Error(err))
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if token != nil {
		claims, ok := token.Claims.(*JWTClaim)
		if !ok {
			zap.L().Error("Validate token error: couldn't parse claims")
			err = errors.New("couldn't parse claims")
			return nil, err
		}

		if claims.ExpiresAt < time.Now().Local().Unix() {
			zap.L().Error("Validate token error: token expired")
			err = errors.New("token expired")
			return nil, err
		}

		if token.Valid && ok {
			zap.L().Info("Validate user " + claims.Username + " 's token success")
			return claims, nil
		}
	} else {
		zap.L().Error("Validate Token error：token invalid")
	}

	zap.L().Error("Validate Token error：unknow error")
	return nil, err

}
