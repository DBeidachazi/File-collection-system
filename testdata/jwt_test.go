package testdata

import (
	"FengfengStudy/core"
	"FengfengStudy/global"
	"FengfengStudy/utils/jwts"
	"testing"
)

func Test(t *testing.T) {
	core.InitConf()
	core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: "好孩子",
		StuId:    20212320,
	})
	if err != nil {
		global.Log.Errorln(token, err)
		return
	}
}
