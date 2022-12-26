package user

type SignUpDTO struct {
	Username  string `db:"username"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	Code      int    `db:"code"`
}

type VerifyUserSt struct {
	Email string `db:"email"`
}
