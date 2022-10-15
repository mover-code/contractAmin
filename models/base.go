package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	orm *gorm.DB
	err error
)

func Init(s, m db.Connection) {
	// fmt.Println("初始化connection", m.GetConfig("main"), s.GetConfig("default"))
	orm, err = gorm.Open(sqlite.Open(s.GetConfig("default").File), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// orm.Use(
	// 	dbresolver.Register(
	// 	dbresolver.Config{
	// 		Replicas: []gorm.Dialector{mysql.Open(m.GetConfig("main").GetDSN())},
	// 	}, "blind").Register(dbresolver.Config{
	// 	Replicas: []gorm.Dialector{mysql.Open(m.GetConfig("exchange").GetDSN())},
	// }, "exchange")
	// )

	if err != nil {
		panic("initialize orm failed")
	}
}
