package dto

type RegisterDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email,min=12" db:"email"`
	Password string `json:"password" form:"password" binding:"required,min=6" db:"password"`
}

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email" db:"email"`
	Password string `json:"password" form:"password" binding:"required" db:"password"`
}
