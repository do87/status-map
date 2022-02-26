package database

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const ERR_CODE_NO_DB = "3D000"
const ERR_CODE_BAD_CRED = "28P01"

func getGormConfig(isLocal bool) *gorm.Config {
	level := logger.Silent
	if isLocal {
		level = logger.Info
	}
	return &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\n", log.LstdFlags),
			logger.Config{LogLevel: level, Colorful: false, SlowThreshold: time.Second},
		),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}
}
