package models

import (
	"fmt"
	"time"
	"yanbao/pkg/logging"
	"yanbao/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var db *gorm.DB

// Setup 初始化连接
func Setup() {
	var dbURI string
	var dialector gorm.Dialector
	dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name)
	dialector = mysql.New(mysql.Config{
		DSN:                       dbURI, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logging.LogErrorWithFields(err.Error(), logging.Fields{})
		panic(fmt.Sprintf("connect db server failed: %v", err.Error()))
	}
	sqlDb, err := conn.DB()
	if err != nil {
		panic(fmt.Sprintf("connect db server failed: %v", err))
	}

	sqlDb.SetMaxIdleConns(setting.DatabaseActivitySetting.MaxIdleConns)
	sqlDb.SetMaxOpenConns(setting.DatabaseActivitySetting.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Second * 600)

	conn.Use(
		dbresolver.Register(dbresolver.Config{
			// `db3`、`db4` 作为 replicas
			Replicas: []gorm.Dialector{mysql.Open("activity"), mysql.Open("demo_0")},
			// sources/replicas 负载均衡策略
			Policy: dbresolver.RandomPolicy{},
		}))
	db = conn
}

// GetDB 开放给外部获得db连接
func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		logging.LogErrorWithFields(fmt.Sprintf("connect db server failed: %v", err), logging.Fields{})
		Setup()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		Setup()
	}

	return db
}
