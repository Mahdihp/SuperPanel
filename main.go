package main

import (
	config "SuperPanel/config"
	"SuperPanel/domain"
	_userController "SuperPanel/user/delivery/http"
	_mysqlUserRepo "SuperPanel/user/repository/mysql"
	_userUseCase "SuperPanel/user/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
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

	app := fiber.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	UserRepo := _mysqlUserRepo.NewMySqlUserRepository(DB)
	UserUseCase := _userUseCase.NewUserUsecase(UserRepo, timeoutContext)
	_userController.NewUserController(app, UserUseCase)
	app.Listen(":3000")
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
