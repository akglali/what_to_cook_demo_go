package user

type signUpDTO struct {
	Username string `db:"username" form:"username" json:"username" validate:"required,min=6,max=18"`
	Password string `db:"password" form:"password" json:"password" validate:"required,min=6,max=18"`
	Email    string `db:"email" form:"email" json:"email" validate:"required"`
	Code     string `db:"code" form:"code" json:"code" validate:"required"`
}

type verifyUserDTO struct {
	Email string `db:"email" form:"email" json:"email" validate:"required"`
}

type googleSignInDTO struct {
	IdToken string `form:"id_token" json:"id_token" validate:"required"`
}
