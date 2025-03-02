package models

import (
	"order-manage/utils/pagination"
	"time"
)

type Customer struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"size:225;not null"`
	Email     string     `gorm:"size:255;not null;unique" json:"email"`
	Phone     string     `gorm:"size:255;not null;unique" json:"phone"`
	Address   string     `gorm:"size:255;not null" json:"address"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt"`
	CreatedBy string     `gorm:"size:225;not null" json:"createdBy"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt"`
	UpdatedBy string     `gorm:"size:225;not null"`
	DeletedAt *time.Time `json:"-"`
	DeletedBy *string    `json:"-" gorm:"size:225"`
}

type ListCustomer struct {
	Data       []Customer            `json:"customers"`
	Pagination pagination.Pagination `json:"meta"`
}

type ParamCustomer struct {
	Limit   uint64 `json:"limit"`
	Page    uint64 `json:"page"`
	Sorting string `json:"sorting"`
	Search  string `json:"search"`
}
