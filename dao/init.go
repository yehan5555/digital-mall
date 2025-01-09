package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	//日志的打印
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		//打印日志信息
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		//线下环境
		ormLogger = logger.Default
	}
	//连接数据库, 下载驱动
	db, err := gorm.Open(mysql.New(mysql.Config{
		// 主数据库
		DSN:                      connRead,
		DefaultStringSize:        256,  //string 类型字段的默认长度
		DisableDatetimePrecision: true, //禁止 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:   true, //重命名索引时不支持，需要将索引先删除在新建，MySQL 5.7 之后的数据库支持
		DontSupportRenameColumn:  true, //change 重命名列时不支持，需要将列先删除在新建，MySQL 8 之后的数据库支持
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名不加复数

		},
	})
	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)                  // 设置连接池
	sqlDB.SetMaxOpenConns(100)                 // 设置连接池最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 设置连接池最大连接生命周期
	_db = db

	// 主从配置
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)},
			Policy:   dbresolver.RandomPolicy{},
		}))
	migration()

}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
