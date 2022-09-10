package db

import (
	"fmt"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

const (
	// DB vars // @todo use config
	dbName = "tasksdb"
	dbHost = "mysql"
	dbUser = "root"
	dbPass = "root"
	dbPort = "3306"

	// Init param vars
	retryTimes = 5
	sleepTime  = 5
)

func NewDatabase() (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info), //@todo make it a config
	}

	connectionDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(gormMysql.Open(connectionDSN), gormConfig)
	if err == nil {
		return db, nil
	}

	for i := 0; i < retryTimes && err != nil; i++ {
		fmt.Println("retrying connecting to database")
		time.Sleep(sleepTime * time.Second)
		db, err = gorm.Open(gormMysql.Open(connectionDSN), gormConfig)
	}

	if err != nil {
		fmt.Errorf("error connecting to database")
		return nil, err
	}

	return db, nil
}
