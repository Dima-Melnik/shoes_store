package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	IGID      string
	Role      string
	CreatedAt time.Time
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type Attribute struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Value string
}

type Product struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	ImageURL   string
	Price      float64
	CategoryID uint
	Category   Category
	Attributes []Attribute `gorm:"many2many:product_attributes"`
	CreatedAt  time.Time
}

type ProductAttributes struct {
	ProductID   uint
	AttributeID uint
}
