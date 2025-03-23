package Models

import "time"

type User struct {
	Id         uint       `json:"id" gorm:"not null;primaryKey"`
	First_name string     `json:"first_name" gorm:"not null;size:255"`
	Last_name  *string    `json:"last_name" gorm:"size:255"`
	Email      string     `json:"email" gorm:"not null;unique;size:255"`
	Phone      string     `json:"phone" gorm:"not null;unique;size:255"`
	Image      string     `json:"image" gorm:"size:500"`
	Password   string     `json:"password" gorm:"not null;size:255"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}
