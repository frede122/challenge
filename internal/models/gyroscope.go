package models

import "time"

type Gyroscope struct {
	ID        uint64    `gorm:"primaryKey" json:"id,omitempty" `
	MAC       string    `gorm:"index;not null" json:"mac" validate:"required"`
	X         float64   `gorm:"not null" json:"x" validate:"required"`
	Y         float64   `gorm:"not null" json:"y" validate:"required"`
	Z         float64   `gorm:"not null" json:"z" validate:"required"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
