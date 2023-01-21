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

type googleIdTokenResp struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Iat           string `json:"iat"`
	Exp           string `json:"exp"`
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
	Typ           string `json:"typ"`
}
