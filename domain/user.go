package domain

import "context"

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string

	Roles     []Role `gorm:"many2many:User_Role;"`
	Active    bool
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
	DeletedAt int64
}

type UserRepository interface {
	GetByID(ctx context.Context, id uint) (User, error)
}
