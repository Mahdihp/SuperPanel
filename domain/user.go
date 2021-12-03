package domain

import "context"

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Username  string
	Password  string
	Roles     []Role `gorm:"many2many:User_Role;"`
	Active    bool
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
	DeletedAt int64
}

type UserRepository interface {
	GetByID(ctx context.Context, id uint) (User, error)
	GetByUsername(ctx context.Context, Username string, Password string) (User, error)
}
type UserUsecase interface {
	GetByID(ctx context.Context, id uint) (User, error)
}
