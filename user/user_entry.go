package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"time"
	"what_to_cook_demo_go/helpers"
)

func SetupSignUp(rg *gin.RouterGroup) {
	rg.POST("", SignUp)
	rg.POST("/Verify", verifyAccount)
}

func SignUp(c *gin.Context) {
	body := SignUpDTO{}
	helpers.AcceptMethod(&body, c)
	// it will be checked in front end too.
	if helpers.ValidateLen(body.Username, 6) || helpers.ValidateLen(body.Password, 6) {
		helpers.NotAcceptableAbort(c, "Password and username must be more than 6 character")
		return
	}

	password, _ := hashPassword(body.Password)
	token := tokenGenerator()
	//use Exec whenever we want to insert update or delete
	//Doing Exec(query) will not use a prepared statement, so lesser TCP calls to the SQL server

	err := insertUser(body.Username, body.FirstName, body.LastName, password, body.Email, token)
	if err != nil {
		fmt.Println(err)
		helpers.BadRequestAbort(c, "User is already exist")
		return
	}

	c.JSON(200, token)
}

func verifyAccount(c *gin.Context) {
	body := VerifyUserSt{}
	helpers.AcceptMethod(&body, c)
	code, _ := GenerateOTP(6)
	email := body.Email
	subject := "Test Email"
	emailBody := "Your code is " + code
	now := time.Now().Format("01-02-2006 15:04:05")
	err := insertVerify(email, code, now)
	if err != nil {
		// if error is because of duplication we update the database. 23505 is duplicate error code
		if string(err.(*pq.Error).Code) == "23505" {
			err = updateVerify(email, code, now)
			err = sendEmail(subject, emailBody, email)
			if err != nil {
				helpers.BadRequestAbort(c, "Code is not sent please try again")
				return
			}
			c.JSON(200, "Your code is sent to "+body.Email)
		} else {
			helpers.BadRequestAbort(c, "Something went wrong!!")
		}
		return
	}
	err = sendEmail(subject, emailBody, email)
	if err != nil {
		helpers.BadRequestAbort(c, "Code is not sent please try again")
		return
	}

	c.JSON(200, "Your code is sent to "+body.Email)

}
