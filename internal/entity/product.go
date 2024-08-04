package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"type:varchar(100)"`
	Rating         float64   `gorm:"type:decimal(2,1)"`
	ReviewsCount   int       `gorm:"type:int"`
	ReviewersCount int       `gorm:"type:int"`
	Description    string    `gorm:"type:text"`
	CurrentPrice   int       `gorm:"type:int"`
	OriginalPrice  int       `gorm:"type:int"`
	Currency       string    `gorm:"type:varchar(10)"`
	Colors         []Color   `gorm:"many2many:product_colors;"`
	Sizes          []Size    `gorm:"many2many:product_sizes;"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
}

type Color struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(50)"`
	Hex  string `gorm:"type:varchar(7)"`
}

type Size struct {
	ID    uint `gorm:"primaryKey"`
	Value int  `gorm:"type:int"`
}

type ProductColor struct {
	ProductID uint `gorm:"primaryKey"`
	ColorID   uint `gorm:"primaryKey"`
}

type ProductSize struct {
	ProductID uint `gorm:"primaryKey"`
	SizeID    uint `gorm:"primaryKey"`
}

func MigrateProductEntityDomain(db *gorm.DB) error {
	err := db.AutoMigrate(&Product{}, &Color{}, &Size{}, &ProductColor{}, &ProductSize{})
	if err != nil {
		return err
	}

	return nil
}
