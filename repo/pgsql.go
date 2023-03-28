package repo

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"grpc-project/common"

	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (*gorm.DB, error) {
	config := &gorm.Config{}
	switch common.Config.Log.Level {
	case "DEBUG", "debug", "5":
		config.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: false,       // 是否忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,        // 是否打开彩色打印
			},
		)
	}

	var db *gorm.DB
	var err error

	switch common.Config.Data.DataBase.Driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(fmt.Sprintf(common.Config.Data.DataBase.Source, common.Config.Data.DataBase.DBName)), config)
	case "mysql":
		db, err = gorm.Open(mysql.Open(common.Config.Data.DataBase.Source), config)
	}
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConn 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConn 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// CreateDatabase 创建数据库，存在则跳过
func CreateDatabase() error {

	defaultDB := common.Config.Data.DataBase.Driver
	conn := fmt.Sprintf(common.Config.Data.DataBase.Source, common.Config.Data.DataBase.DefaultDBName)

	db, err := sql.Open(defaultDB, conn)
	if err != nil {
		log.Default().Println("open sql failed:", err.Error())
		return err
	}

	defer db.Close()

	exec := fmt.Sprintf("CREATE DATABASE \"%v\" WITH ENCODING = 'utf8'", common.Config.Data.DataBase.DBName)
	_, err = db.Exec(exec)

	if err == nil {
		log.Default().Printf("successfully create database %s\n", common.Config.Data.DataBase.DBName)
		return nil
	}

	if err != nil && strings.Contains(err.Error(), "already exists") {
		log.Default().Printf("database %s already exists，continue\n", common.Config.Data.DataBase.DBName)
		return nil
	}
	log.Default().Printf("fail create database: : %v", common.Config.Data.DataBase.DBName)
	return err

}
