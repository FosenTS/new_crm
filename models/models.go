package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"` // "employee" or "boss"
}

type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Quantity    int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`
}

type Order struct {
	gorm.Model
	UserID      uint    `gorm:"not null"`
	User        User    `gorm:"foreignKey:UserID"`
	Status      string  `gorm:"not null"` // "pending", "processing", "completed", "cancelled"
	TotalAmount float64 `gorm:"not null"`
	Items       []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
}
