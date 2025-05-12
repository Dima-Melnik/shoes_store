package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	IGID      string    `json:"idig"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`
}

type Attribute struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Product struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	Name       string      `json:"name"`
	ImageURL   string      `json:"image_url"`
	Price      float64     `json:"price"`
	CategoryID uint        `json:"category_id"`
	Category   Category    `json:"category"`
	Attributes []Attribute `json:"attributes" gorm:"many2many:product_attributes"`
	CreatedAt  time.Time   `json:"created_at"`
}

type ProductAttributes struct {
	ProductID   uint
	AttributeID uint
}

type CreateProductInput struct {
	Name         string  `json:"name"`
	ImageURL     string  `json:"image_url"`
	Price        float64 `json:"price"`
	CategoryID   uint    `json:"category_id"`
	AttributeIDs []uint  `json:"attribute_ids"`
}
