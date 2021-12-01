package domain

type Role struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string

	Permissions []Permission `gorm:"many2many:role_permission;"`
	Active      bool
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UpdatedAt   int64 `gorm:"autoUpdateTime:milli"`
	DeletedAt   int64
}
