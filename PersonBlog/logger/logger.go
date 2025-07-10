package logger

import (
	"PersonBlog/config"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	LodCode uint
	LogName string
	Err     interface{}
}

func InitLog() {
	err := config.DB.AutoMigrate(&Log{})
	if err != nil {
		return
	}
}

func AddLog(code uint, logName string, err interface{}) {
	log := &Log{
		LodCode: code,
		LogName: logName,
		Err:     err,
	}
	config.DB.Model(&Log{}).Create(log)
}
