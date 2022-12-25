package user

import "what_to_cook_demo_go/database"

func insertVerify(email, code, now string) error {
	_, err := database.Db.Exec("insert into verify_users_table(email,code,created_at) values ($1,$2,$3)", email, code, now)
	if err != nil {
		return err
	}
	return nil
}

func updateVerify(email, code, now string) error {
	_, err := database.Db.Exec("update verify_users_table set code=$1 , created_at=$2 where email=$3", code, now, email)
	if err != nil {
		return err
	}
	return nil
}

func insertUser(username, firstName, lastName, password, email, token string) error {
	_, err := database.Db.Exec("insert into users_table(username,firstname,lastname,password,email,token) values ($1,$2,$3,$4,$5,$6)", username, firstName, lastName, password, email, token)
	if err != nil {
		return err
	}
	return nil
}
