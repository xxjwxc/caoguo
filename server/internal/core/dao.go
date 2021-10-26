package core

import (
	"caoguo/internal/config"
	"runtime"
	"time"

	"github.com/xxjwxc/public/errors"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/mysqldb"
)

// Dao ...
var Dao DaoCore

// DaoCore core dao
type DaoCore struct {
	dbr *mysqldb.MySqlDB // 读库
	dbw *mysqldb.MySqlDB // 写库
}

func init() {
	Dao.InitDao()
}

// GetDB 获取读数据库
func (dao *DaoCore) GetDB() *mysqldb.MySqlDB {
	if dao.dbr == nil {
		mylog.Error(errors.New("dbr is nil. "))
	}
	return dao.dbr
}

// GetDBr 获取读数据库
func (dao *DaoCore) GetDBr() *mysqldb.MySqlDB {
	if dao.dbr == nil {
		mylog.Error(errors.New("dbr is nil. "))
	}
	return dao.dbr
}

// GetDBw 获取写数据库
func (dao *DaoCore) GetDBw() *mysqldb.MySqlDB {
	if dao.dbw == nil {
		mylog.Error(errors.New("dbw is nil. "))
	}
	return dao.dbw
}

// InitDao 初始化dao
func (dao *DaoCore) InitDao() {
	runtime.SetFinalizer(dao, dao.Destroy) //析构时触发

	dao.dbr = mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	sqlDB, _ := dao.dbr.DB.DB()
	sqlDB.SetMaxIdleConns(10)           // 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxOpenConns(100)          // 设置数据库的最大连接数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接的最大可复用时间。
	dao.dbw = mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	sqlDB, _ = dao.dbw.DB.DB()
	sqlDB.SetMaxIdleConns(10)           // 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxOpenConns(100)          // 设置数据库的最大连接数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接的最大可复用时间。
}

// Destroy 释放
func (dao *DaoCore) Destroy() {
	if dao.dbr != nil {
		dao.dbr.OnDestoryDB()
	}

	if dao.dbw != nil {
		dao.dbw.OnDestoryDB()
	}
}
