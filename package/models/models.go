package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"not null" json:"username"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Gender   string `json:"gender"`
	Phone    int    `gorm:"not null" json:"phone"`
	Blocking bool   `json:"blocking"`
}

type OtpMail struct {
	Id        uint
	Email     string `gorm:"unique" json:"email"`
	Otp       string `gorm:"not null" json:"otp"`
	CreatedAt time.Time
	ExpireAt  time.Time `gorm:"type:timestamp;not null"`
}

type Products struct {
	gorm.Model
	Name        string `gorm:"unique" json:"p_name"`
	Price       uint   `json:"p_price"`
	Size        string `json:"p_size"`
	Color       string `json:"p_color"`
	Quantity    int    `json:"p_quantity"`
	Description string `json:"p_description"`
	ImagePath1  string
	ImagePath2  string
	ImagePath3  string
	Status      bool `json:"p_blocking"`
	CategoryId  int  `json:"category_id"`
	Category    Category
}
type Category struct {
	gorm.Model
	Category_name        string `gorm:"not null" json:"category_name"`
	Category_description string `gorm:"not null" json:"category_description"`
	Blocking             bool   `gorm:"not null" json:"category_blocking"`
}
type Rating struct {
	gorm.Model
	Users     int `json:"rating_user"`
	ProductId int `gorm:"unique" json:"rating_product"`
	Product   Products
	Value     int `json:"rating_value"`
}
type Review struct {
	Id        uint
	UserId    int `json:"review_user"`
	User      Users
	ProductId uint `json:"review_product"`
	Product   Products
	Review    string `json:"review"`
	Time      string
}
