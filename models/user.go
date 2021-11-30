package models

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string

	Active    bool
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
	DeletedAt int64
}
