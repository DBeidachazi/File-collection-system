package global

import (
	"FengfengStudy/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// Config 用于保存配置文件
	// todo 尝试封装查询方法 https://gorm.io/zh_CN/gen/dao.html
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
