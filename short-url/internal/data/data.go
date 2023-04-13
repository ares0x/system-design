package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"short-url/internal/conf"
	"time"
)

var (
	Orm *gorm.DB
)

func Init(config *conf.ServerConfig) (err error) {
	Orm, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DataBase.User,
		config.DataBase.Pwd,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.Name))
	if err != nil {
		err = fmt.Errorf("db connect fail %s", err)
		return
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return fmt.Sprintf("%s%s", config.DataBase.Pref, defaultTableName)
	}
	// 连接池
	Orm.DB().SetMaxIdleConns(50)
	Orm.DB().SetMaxOpenConns(300)
	Orm.DB().SetConnMaxLifetime(300 * time.Second)
	// 全局禁用表名复数
	Orm.SingularTable(true)
	// 禁止没有条件的更新/删除
	Orm.BlockGlobalUpdate(true)
	return
}
