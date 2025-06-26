package Models

import (
	"time"
)

type Session struct {
	Id           uint       `json:"id" gorm:"not null;primaryKey"`
	Token_string string     `json:"token_string"`
	User_id      uint       `json:"user_id" gorm:"not null;primaryKey"`
	User         *User      `json:"user" gorm:"foreignKey:User_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Created_at   *time.Time `json:"created_at"`
	Updated_at   *time.Time `json:"updated_at"`
}
