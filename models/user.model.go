package models

type User struct {
	Name     string   `json:"name" bson:"name" validate:"required"`
	Email    string   `json:"email" bson:"email" validate:"required"`
	Password string   `json:"password" bson:"password" validate:"required"`
	FormIds  []string `json:"form_ids" bson:"form_ids"`
}
