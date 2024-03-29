package models

import "time"

type User struct {
	ID        uint      `gorm:"PrimaryKey"`
	Email     string    `gorm:"not null;unique;type:varchar(191)"`
	Products  []Product `gorm:"foreignKey:UserID`
	CreatedAt time.Time
	UpdatedAt time.Time
}
