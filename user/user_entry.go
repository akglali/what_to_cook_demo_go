package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"time"
	"what_to_cook_demo_go/helpers"
)

// SetupEntry it is a prefix routerGroup for entries
func SetupEntry(rg *gin.RouterGroup) {
	rg.POST("", SignUp)
	rg.POST("/Verify", VerifyAccount)
}

func SignUp(c *gin.Context) {
	body := SignUpDTO{}
	passedMarshall := helpers.AcceptMethod(&body, c)
	if !passedMarshall {
		return
	}
	var createdAt string
	var code string
	if err := verifyUser(body.Email, &createdAt, &code); err != nil {
		fmt.Println(err)
		helpers.NotAcceptableAbort(c, "Something went wrong!")
		return
	}
	//check if user enters the code within 121 seconds
	if isExpired := timeDiffSec(createdAt); isExpired == false {
		helpers.NotAcceptableAbort(c, "Time is expired send another code!")
		return
	}
	//check if the code is true
	if body.Code != code {
		helpers.NotAcceptableAbort(c, "Code is not true!")
		return
	}

	password, _ := hashPassword(body.Password)
	token := tokenGenerator()
	if err := insertUser(body.Username, password, body.Email, token); err != nil {
		fmt.Println(err)
		helpers.BadRequestAbort(c, "User is already exist")
		return
	}
	userToken, err := helpers.Encrypt(token)
	if err != nil {
		fmt.Println(err)
		helpers.BadRequestAbort(c, "Something went wrong")
		return
	}
	c.JSON(200, userToken)
}

func VerifyAccount(c *gin.Context) {
	body := VerifyUserDTO{}
	passedMarshall := helpers.AcceptMethod(&body, c)
	if !passedMarshall {
		return
	}
	//checking if user is already exist before sent the email.
	err, isUserExist := userExist(body.Email)
	if err != nil || isUserExist {
		fmt.Println(err)
		helpers.BadRequestAbort(c, "User is already exist")
		return
	}
	code, _ := GenerateOTP(6)
	email := body.Email
	subject := "Test Email"
	emailBody := "Your code is " + code
	now := time.Now().Format(dateFormat)
	if err := insertVerify(email, code, now); err != nil {
		// if error is because of duplication we update the database. 23505 is duplicate error code
		if string(err.(*pq.Error).Code) == "23505" {
			//fmt.Println("user hasn't verified yet")
			if err := updateVerify(email, code, now); err != nil {
				// if database can't update the code.
				helpers.BadRequestAbort(c, "Something went wrong")
				return
			}
			//send email to the same address since it is already created.
			if err := sendEmail(subject, emailBody, email); err != nil {
				helpers.BadRequestAbort(c, "Code is not sent please try again")
				return
			}
			c.JSON(200, "Your code is sent to "+body.Email)
			return
		} else {
			helpers.BadRequestAbort(c, "Something went wrong!!")
		}
		return
	}
	//that sends the email
	if err := sendEmail(subject, emailBody, email); err != nil {
		helpers.BadRequestAbort(c, "Code is not sent please try again")
		return
	}
	c.JSON(200, "Your code is sent to "+body.Email)

}

/*func signIn() {

}
*/
/*_, err, actualToken := helpers.UserAuth(userToken)
if err != nil {
fmt.Println(err.Error())
}
fmt.Println(actualToken)*/
