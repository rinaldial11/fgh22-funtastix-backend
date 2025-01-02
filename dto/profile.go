package dto

type ProfileDTO struct {
	FirstName       string `form:"first_name" binding:"min=3"`
	LastName        string `form:"last_name" binding:"min=3"`
	Email           string `form:"email" binding:"required,email"`
	PhoneNumber     string `form:"phone_number"`
	Picture         string `form:"-" swaggerignore:"true"`
	Point           string `form:"point"`
	Password        string `form:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`
}
