package dao

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"project/configs"
)
var Provider = wire.NewSet( NewDB)
type Dao struct {
	db *gorm.DB
}

func NewDB() (*Dao, func(), error) {
	// TODO 单例封装
	// TODO 配置Config
	connString := configs.ProjectConfig.MysqlUrl
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return nil, nil, fmt.Errorf("mysql conn err,err is:%v", err)
	}
	closeFunc := func() { db.Close() }
	return &Dao{
		db: db,
	}, closeFunc, nil
}

