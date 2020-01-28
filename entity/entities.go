package entity

import "time"

// Gallery represents Food Menu Gallery
type Gallery struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Image       string `gorm:"type:varchar(255)"`
	Pic         []Pic  `gorm:"many2many:pic_gallerys"`
}

// User represents application user
type User struct {
	ID       uint
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
}

// Role repesents application user roles
type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}

// Item represents food menu pics
type Pic struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Gallerys  []Gallery `gorm:"many2many:pic_gallerys"`
	Image       string     `gorm:"type:varchar(255)"`
}

//Session represents login user session
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

// Comment represents comments forwarded by application users
type Comment struct {
	ID        uint
	FullName  string `gorm:"type:varchar(255)"`
	Message   string
	Email     string `gorm:"type:varchar(255);not null; unique"`
	CreatedAt time.Time
}
