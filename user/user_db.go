package user

import (
	"what_to_cook_demo_go/database"
)

// **** Use Exec whenever we want to insert update or delete
// **** Doing Exec(query) will not use a prepared statement, so lesser TCP calls to the SQL server

// this code will save the code that has been sent to user's email
func insertVerify(email, code, now string) error {
	_, err := database.Db.Exec("insert into verify_users_table(email,code,created_at) values ($1,$2,$3)", email, code, now)
	if err != nil {
		return err
	}
	return nil
}

// if user sent the code again it will update the code in the system
func updateVerify(email, code, now string) error {
	_, err := database.Db.Exec("update verify_users_table set code=$1 , created_at=$2 where email=$3", code, now, email)
	if err != nil {
		return err
	}
	return nil
}

// it inserts the user if user has passed all the required action successfully
func insertUser(username, password, email, token string) error {
	_, err := database.Db.Exec("insert into users_table(username,password,email,token) values ($1,$2,$3,$4)", username, password, email, token)
	if err != nil {
		return err
	}
	// since user is already verified we can delete it from verify user table
	deleteUserCode(email)
	return nil
}

// user code is removed from the database
func deleteUserCode(email string) {
	_, err := database.Db.Exec("delete from verify_users_table where email=$1", email)
	if err != nil {
		return
	}
}

// verify the users code and expiration
func verifyUser(email string, createdAt *string, code *int) error {
	err := database.Db.QueryRow("select created_at,code from verify_users_table where email=$1", email).Scan(&*createdAt, &*code)
	if err != nil {
		return err
	}
	return nil
}

// if user is already exist we don't send a code
func userExist(email string) (error, bool) {
	var isUserExist bool
	err := database.Db.QueryRow("select exists(select 1 from users_table where email=$1)", email).Scan(&isUserExist)
	if err != nil {
		return err, false
	}
	if isUserExist == true {
		return nil, true
	}
	return nil, false
}
