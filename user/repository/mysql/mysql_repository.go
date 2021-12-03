package mysql

import (
	"SuperPanel/domain"
	"context"
	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	DB *gorm.DB
}

func NewMySqlUserRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepo{
		DB: db,
	}
}
func (this *mysqlUserRepo) GetByID(ctx context.Context, id int32) (domain.User, error) {
	var user domain.User
	first := this.DB.First(&user, id)
	if first.Error == nil {
		return user, nil
	}
	return domain.User{}, first.Error
}

func (this *mysqlUserRepo) GetByUsername(ctx context.Context, Username string, Password string) (domain.User, error) {
	var user domain.User
	find := this.DB.Where(map[string]interface{}{"Username": Username, "Password": Password}).Find(&user)
	if find.Error == nil {
		return user, nil
	}
	return domain.User{}, find.Error
}
