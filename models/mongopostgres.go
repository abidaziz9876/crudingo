package models

import "time"

type MongoPostgres struct {
	ID         int64     `gorm:"column:id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	Name       string    `gorm:"column:name"`
	Age        int       `gorm:"column:age"`
	Salary     int       `gorm:"salary"`
	Department string    `gorm:"department"`
}
