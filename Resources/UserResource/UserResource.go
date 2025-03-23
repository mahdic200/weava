package UserResource

import (
	"time"

	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Utils"
)

type User struct {
	Id         uint       `json:"id"`
	First_name string     `json:"first_name"`
	Last_name  *string    `json:"last_name"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Image      string     `json:"image"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated"`
}

func Single(user Models.User) User {
	single := User{
		Id:         user.Id,
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Email:      user.Email,
		Phone:      user.Phone,
		Image:      Utils.ImageUrlOrDefault(user.Image),
		Created_at: user.Created_at,
		Updated_at: user.Updated_at,
	}
	return single
}

func Collection(users []Models.User) []User {
	collection := []User{}
	for _, user := range users {
		single := Single(user)
		collection = append(collection, single)
	}
	return collection
}
