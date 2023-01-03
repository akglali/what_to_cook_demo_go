package user

type SignUpDTO struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Code     int    `db:"code"`
}

type VerifyUserSt struct {
	Email string `db:"email"`
}
