package dto

type AuthDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email" db:"email"`
	Password string `json:"password" form:"password" binding:"required,min=5" db:"password"`
}
