package mysql

import (
	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	DB *gorm.DB
}
