package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type App struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type MySQLInfo struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type Config struct {
	App   App       `mapstructure:"app"`
	MySQL MySQLInfo `mapstructure:"mysql"`
}

var Conf *Config
var DB *gorm.DB

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigName("config") // 不带扩展名
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 当前目录

	if err := viper.ReadInConfig(); err != nil {
		panic("读取配置文件失败: " + err.Error())
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic("解析配置文件失败: " + err.Error())
	}

	println("配置加载成功")

	mysqlConf := Conf.MySQL

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database,
		mysqlConf.Charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加 s，user 而不是 users
		},
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	DB = db
}
