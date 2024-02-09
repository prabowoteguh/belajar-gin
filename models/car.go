package models

import (
	"time"
)

type Car struct {
	ID        uint   `gorm:"PrimaryKey"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	Model     string `gorm:"not null;type:varchar(191)"`
	Price     int    `gorm:"not null;type:int"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
// 	fmt.Println("Product before create()")

// 	if len(p.Name) < 4 {
// 		err = errors.New("product name is too short")
// 	}

// 	return
// }
