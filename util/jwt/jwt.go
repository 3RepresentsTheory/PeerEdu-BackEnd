package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 用于验证用户登录状态

type UserClaims struct {
	Id     string `json:"id"`
	Expire int64  `json:"exp"`
	Role   uint   `json:"role"` // db.go中的定义
	jwt.RegisteredClaims
}

func GetUserToken(id string, expire int64, key string, identity uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		id,
		time.Now().Unix() + expire,
		identity,
		jwt.RegisteredClaims{},
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// 验证用户 返回用户id以及身份
func VerifyUserToken(tokenString string, key string) (string, uint) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	if err != nil {
		println(err.Error())
		return "", 0
	}

	if claims, ok := token.Claims.(*UserClaims); ok &&
		token.Valid && claims.Expire > time.Now().Unix() {
		return claims.Id, claims.Role
	} else {
		return "", 0
	}
}
