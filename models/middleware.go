package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"practice_go/config"
	"time"
)

type Model struct {
	ID uint 	`gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Item struct {
	Model
	ItemName string `gorm:"size:255" json:"item_name,omitempty"`
	Price int `json:"price,omitempty"`
}