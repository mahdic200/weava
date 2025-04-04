package Models

import (
	"time"
)

type AdminSession struct {
	Id         uint       `json:"id" gorm:"not null;primaryKey"`
	Admin_id   uint       `json:"admin_id" gorm:"not null;primaryKey"`
	Admin      *Admin     `json:"admin" gorm:"foreignKey:Admin_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}
