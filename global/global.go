package global

import (
	"FengfengStudy/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// 用于保存配置文件
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
