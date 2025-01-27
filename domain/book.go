package domain

import "time"

type Book struct {
	Id          string  `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Price       float64 `gorm:"type:decimal(10,2);default:0.00;not null"`
	PublishDate time.Time
	Uid         string `gorm:"index"`
}
