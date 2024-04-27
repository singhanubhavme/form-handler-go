package models

type Form struct {
	Email       string `json:"email" bson:"email" validate:"required"`
	FormTitle   string `json:"form_title" bson:"email" validate:"required"`
	FormId      string `json:"form_id" bson:"form_id"`
	RedirectUrl string `json:"redirect_url" bson:"redirect_url" validate:"required"`
}
