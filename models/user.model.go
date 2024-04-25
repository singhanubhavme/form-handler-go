package models

type User struct {
	Name     string   `json:"name" validate:"required"`
	Email    string   `json:"email" validate:"required"`
	Password string   `json:"password" validate:"required"`
	FormIds  []string `json:"form_ids"`
}
