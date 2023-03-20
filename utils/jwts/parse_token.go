package jwts

import (
	"FengfengStudy/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
)

// ParseToken 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	Mysecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Mysecret, nil
	})
	if err != nil {
		global.Log.Warnln(fmt.Sprintf("token解析失败: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}
