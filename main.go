package main

import (
	config2 "SuperPanel/config"
	"SuperPanel/models"
	"fmt"
	"github.com/spf13/viper"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var config config2.Configurations

func init() {
	InitConfig()
}
func main() {
	db := InitDb()
	fmt.Println(db.Config.Name())
}

func InitDb() *gorm.DB {
	dsn := config.Database.DBUser + ":" +
		config.Database.DBPassword + "@tcp(127.0.0.1:" +
		strconv.Itoa(config.Database.DBPort) + ")/" + config.Database.DBName +
		"?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Permission{})
	return db
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetDefault("database.dbname", "super_panel")
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Database is\t", config)
}
