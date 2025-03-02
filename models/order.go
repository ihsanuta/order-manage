package models

import (
	"order-manage/utils/pagination"
	"time"
)

type Order struct {
	ID          uint64     `gorm:"primary_key" json:"id"`
	CustomerID  uint64     `gorm:"not null" json:"customerID"`
	Description string     `gorm:"size:225;not null"`
	Total       uint64     `gorm:"not null" json:"total"`
	CreatedAt   time.Time  `gorm:"not null" json:"createdAt"`
	CreatedBy   string     `gorm:"size:225;not null" json:"createdBy"`
	UpdatedAt   time.Time  `gorm:"not null" json:"updatedAt"`
	UpdatedBy   string     `gorm:"size:225;not null"`
	DeletedAt   *time.Time `json:"-"`
	DeletedBy   *string    `json:"-" gorm:"size:225"`
}

type ListOrder struct {
	Data       []Order               `json:"orders"`
	Pagination pagination.Pagination `json:"meta"`
}

type ParamOrder struct {
	Limit   uint64 `json:"limit"`
	Page    uint64 `json:"page"`
	Sorting string `json:"sorting"`
	Search  string `json:"search"`
}
