package user

type SignUpDTO struct {
	Username string `db:"username" form:"username" json:"username" validate:"required,min=6,max=18"`
	Password string `db:"password" form:"password" json:"password" validate:"required,min=6,max=18"`
	Email    string `db:"email" form:"email" json:"email" validate:"required"`
	Code     string `db:"code" form:"code" json:"code" validate:"required"`
}

type VerifyUserDTO struct {
	Email string `db:"email" form:"email" json:"email" validate:"required"`
}
