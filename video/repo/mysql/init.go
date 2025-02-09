package mysql

import (
	"fmt"
	"time"

	"github.com/Yamon955/ShortVideo/video/entity/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

// db 共享的数据库连接对象
var db *gorm.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetAppConfig().MySQLConf.User, conf.GetAppConfig().MySQLConf.Pwd, conf.GetAppConfig().MySQLConf.Address,
		conf.GetAppConfig().MySQLConf.Port, conf.GetAppConfig().MySQLConf.DBName)
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Errorf("gorm.Open with dsn[%s] failed, err:%v", dsn, err)
		return err
	}

	// GORM 使用 database/sql 来维护连接池
	sqlDB, _ := db.DB()
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	//// AutoMigrate 根据 User 结构体，自动创建表结构, 可以通过Set设置附加参数
	//if err := db.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&base.User{}); err != nil {
	//	log.Errorf("db.Migrate failed, err:%v", err)
	//	return err
	//}
	//// AutoMigrate 根据 Video 结构体，自动创建表结构, 可以通过Set设置附加参数
	//if err := db.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&base.Video{}); err != nil {
	//	log.Errorf("db.Migrate failed, err:%v", err)
	//	return err
	//}
	//// AutoMigrate 根据 User 结构体，自动创建表结构, 可以通过Set设置附加参数
	//if err := db.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&base.Like{}); err != nil {
	//	log.Errorf("db.Migrate failed, err:%v", err)
	//	return err
	//}
	//// AutoMigrate 根据 User 结构体，自动创建表结构, 可以通过Set设置附加参数
	//if err := db.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&base.Collect{}); err != nil {
	//	log.Errorf("db.Migrate failed, err:%v", err)
	//	return err
	//}

	return nil
}
