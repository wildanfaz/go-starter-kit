package models

import "time"

type (
	User struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		IsDeleted bool      `json:"is_deleted"`
	}
)
