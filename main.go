package main

import (
	config "SuperPanel/config"
	"SuperPanel/domain"
	"go.uber.org/zap"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var Config *config.Configurations
var Logger *zap.Logger

func init() {
	Config = config.InitConfig()
	Logger = config.InitLog()
}

func main() {
	DB := InitDb()
	Logger.Info(DB.Config.Name())
	Logger.Info(Config.Database.DBName)

}

func InitDb() *gorm.DB {
	dsn := Config.Database.DBUser + ":" +
		Config.Database.DBPassword + "@tcp(127.0.0.1:" +
		strconv.Itoa(Config.Database.DBPort) + ")/" + Config.Database.DBName +
		"?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Role{})
	db.AutoMigrate(&domain.Permission{})
	return db
}
