package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

// 定义JWT的过期时间，这里以2小时
const TokenExpireDuration = time.Hour * 2000

var MySecret = []byte("abcd1234567890")

// GenToken 生成JWT
func GenToken(UserId int) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserId, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "go-api",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, ParseTokenError
}
