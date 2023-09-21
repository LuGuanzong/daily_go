package dao

import (
	"daily_plan_go/common/logger"
	"daily_plan_go/config"
	"daily_plan_go/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	Db  *gorm.DB
	log *logrus.Logger
	err error
)

func init() {
	log, _ = logger.InitLogger()
}

func InitDb(env string) error {
	switch env {
	case "dev":
		err = sqlite()
	case "prod":
		err = mysql()
	default:
		err = sqlite()
	}
	if err != nil {
		return err
	}

	if err = autoMigrate(); err != nil {
		return err
	}

	return nil
}

func autoMigrate() error {
	err = Db.AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Errorf("数据库迁移错误 err: %s", err.Error())
		return err
	}
	return nil
}

func sqlite() error {
	Db, err = gorm.Open("sqlite3", config.Conf.Sqlite.Pwd)
	if err != nil {
		log.Errorf("初始化sqlite出错 err: %s", err.Error())
		return err
	}
	if Db.Error != nil {
		log.Errorf("数据库sqlite出错 err: %s", err.Error())
		return Db.Error
	}

	return nil
}

func mysql() error {
	Db, err = gorm.Open("sqlite", config.Conf.Sqlite.Pwd)
	if err != nil {
		// 打log
	}
	if Db.Error != nil {
		// 打log
	}

	// ----------------------- 连接池设置 -----------------------
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	Db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Db.DB().SetConnMaxLifetime(time.Hour)

	return nil
}