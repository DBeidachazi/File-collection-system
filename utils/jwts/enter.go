package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad jwt中的payload数据
type JwtPayLoad struct {
	Username string `json:"username"` // 用户名
	StuId    int32  `json:"stu_id"`   // 学号
}

var Mysecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
