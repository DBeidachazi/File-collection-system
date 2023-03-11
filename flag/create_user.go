package flag

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/orm/model"
	"FengfengStudy/utils/pwd"
	"fmt"
)

func CreateUser() {
	// 创建用户的逻辑
	// 用户名 密码
	var (
		stuId    int32
		name     string
		password string
	)
	fmt.Printf("请输入用户名: ")
	fmt.Scan(&stuId)
	fmt.Printf("请输入姓名: ")
	fmt.Scan(&name)
	fmt.Printf("请输入密码: ")
	fmt.Scan(&password)

	_, err := dal.User.Where(dal.User.StuID.Eq(stuId)).First()
	if err == nil {
		global.Log.Warnln("创建用户失败, 用户名已存在", stuId, password)
	} else {
		// 对密码进行hash
		hashPwd := pwd.HashPwd(password)

		// 入库
		user := model.User{
			StuID:    stuId,
			Name:     name,
			Password: hashPwd,
		}
		err = dal.User.Create(&user)
		if err != nil {
			global.Log.Warnln("入库失败", stuId, password)
		}
	}
}
