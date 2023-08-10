package model_bank

import (
	"time"
)

type Bank struct {
	ID int `gorm:"type:int;primary_key" json:"id,omitempty"`
	Name string	`gorm:"type:varchar(255)" json:"name,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

type CreateBank struct {
	Name string `json:"name" validate:"required"`
}
