package gorm

import (
	"fmt"

	"gin_one/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

func Setup() {

	// 配置MySQL连接参数
	username := setting.DatabaseSetting.User     // 账号
	password := setting.DatabaseSetting.Password // 密码
	host := setting.DatabaseSetting.Host         // 数据库地址，可以是Ip或者域名
	port := setting.DatabaseSetting.Port         // 数据库端口
	Dbname := setting.DatabaseSetting.DbName     // 数据库名
	timeout := setting.DatabaseSetting.Timeout   // 连接超时，10秒

	var err error

	// 拼接下dsn参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。

	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// 设置数据库连接池参数
	sqlDB, err := _db.DB()
	sqlDB.SetMaxOpenConns(setting.DatabaseSetting.MaxConn)     // 设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConn) // 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

// 获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
// 不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}
