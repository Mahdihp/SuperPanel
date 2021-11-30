package models

type Role struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
	DeletedAt int64
}
