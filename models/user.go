package models

import "time"

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Username  string     `gorm:"size:255;not null;unique" json:"username"`
	Password  string     `gorm:"size:255;not null;" json:"-"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt"`
	CreatedBy string     `gorm:"size:225;not null" json:"createdBy"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt"`
	UpdatedBy string     `gorm:"size:225;not null"`
	DeletedAt *time.Time `json:"-" `
	DeletedBy *string    `json:"-" gorm:"size:225"`
}
