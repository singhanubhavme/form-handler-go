package models

type Form struct {
	Email       string `json:"email" validate:"required"`
	FormTitle   string `json:"form_title" validate:"required"`
	FormId      string `json:"form_id"`
	RedirectUrl string `json:"redirect_url" validate:"required"`
}
