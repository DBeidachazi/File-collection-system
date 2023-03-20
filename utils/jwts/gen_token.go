package jwts

import (
	"FengfengStudy/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

func GenToken(user JwtPayLoad) (string, error) {
	Mysecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{ // payload
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 默认2小时过期
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(Mysecret)
}
