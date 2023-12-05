package jwts

import (
	"errors"
	"fmt"
	"kubez_project/config"
	"kubez_project/utils/logs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Username string
	jwt.RegisteredClaims
}

func GenToken(username string) (token string, err error) {
	claim := MyCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                      // 签发者
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(config.JwtSecret))
	logs.Debug(nil, fmt.Sprint("生成jwt-token: ", token))
	return

}

// 传入jwt-token,返回claims对象和error
func ParseToken(token string) (*MyCustomClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil // 此处填写用于解析token的secret
	})
	if err != nil {
		logs.Error(nil, "jwt-token解析失败")
		return nil, err
	}
	if claims, ok := tokenStr.Claims.(*MyCustomClaims); tokenStr.Valid && ok {
		return claims, nil
	} else {
		logs.Error(nil, "无效的jwt-token")
		return nil, errors.New("invalid token")
	}
}
