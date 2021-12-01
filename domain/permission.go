package domain

type Permission struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string

	Active    bool
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
	DeletedAt int64
}
