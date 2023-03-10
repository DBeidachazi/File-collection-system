package core

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

func dsn() string {
	return global.Config.Mysql.Dsn()
}
func BuildGen() *gorm.DB {
	// gorm 连接数据库
	log.Println(dsn())
	db, err := gorm.Open(mysql.Open(dsn()))
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql 连接失败", dsn()))
	}

	// gen 配置
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./models/orm/dal",
		ModelPkgPath: "./models/orm/model",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})
	// 使用数据库
	g.UseDB(db)
	// 生成所有表的DAO接口 配置文件修改
	if global.Config.CreateDB.IsCreate {
		g.ApplyBasic(g.GenerateAllTable()...)
	}
	g.Execute()
	//log.Println("<-----------gen success----------->")
	logrus.Info("gen init success")
	dal.SetDefault(db)
	return db
}
