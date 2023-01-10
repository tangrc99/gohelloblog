package db

import (
	"fmt"
	"github.com/tangrc99/gohelloblog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL = gorm.DB

func NewMySQLFrom(setting *setting.MySQLSetting) *gorm.DB {

	return NewMySQL(setting.User, setting.Password, setting.Url, setting.DBName, setting.Timeout.String())
}

func NewMySQL(usr, pwd, url, database string, timeout ...string) *gorm.DB {

	if len(timeout) == 0 {
		timeout[0] = "10s"
	}
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", usr, pwd, url, database, timeout[0])

	//连接MySQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	return db
}
