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
	rg.POST("/Verify", verifyAccount)
	rg.GET("/something", something)
}

func SignUp(c *gin.Context) {
	body := SignUpDTO{}
	helpers.AcceptMethod(&body, c)
	// it will be checked in front end too.
	if helpers.ValidateLen(body.Username, 6) || helpers.ValidateLen(body.Password, 6) {
		helpers.NotAcceptableAbort(c, "Password and username must be more than 6 character")
		return
	}
	var createdAt string
	var code int
	err := verifyUser(body.Email, &createdAt, &code)
	if err != nil {
		fmt.Println(err)
		helpers.NotAcceptableAbort(c, "Something went wrong!")
		return
	}
	//check if user enters the code within 121 seconds
	isExpired := timeDiffSec(createdAt)
	if isExpired == false {
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
	err = insertUser(body.Username, password, body.Email, token)
	if err != nil {
		fmt.Println(err)
		helpers.BadRequestAbort(c, "User is already exist")
		return
	}
	jwtToken, err := helpers.GenerateJWT(token)
	if err != nil {
		fmt.Println(err)
		helpers.BadRequestAbort(c, "Something went wrong")
		return
	}
	c.JSON(200, jwtToken)
}

func verifyAccount(c *gin.Context) {
	body := VerifyUserSt{}
	helpers.AcceptMethod(&body, c)
	//checking if user is already exist before sent the email.
	err, isUserExist := userExist(body.Email)
	if err != nil || isUserExist {
		fmt.Println(err)
		helpers.BadRequestAbort(c, "User is already exist")
		return
	}
	code, _ := GenerateOTP(6)
	email := body.Email
	/*	subject := "Test Email"
		emailBody := "Your code is " + code*/
	fmt.Println(code)
	now := time.Now().Format(dateFormat)
	err = insertVerify(email, code, now)
	if err != nil {
		// if error is because of duplication we update the database. 23505 is duplicate error code
		if string(err.(*pq.Error).Code) == "23505" {
			fmt.Println("user hasn't verified yet")
			err = updateVerify(email, code, now)
			/*	err = sendEmail(subject, emailBody, email)
				if err != nil {
					helpers.BadRequestAbort(c, "Code is not sent please try again")
					return
				}*/
			c.JSON(200, "Your code is sent to "+body.Email)
			return
		} else {
			helpers.BadRequestAbort(c, "Something went wrong!!")
		}
		return
	}
	//that sends the email
	/*err = sendEmail(subject, emailBody, email)
	if err != nil {
		helpers.BadRequestAbort(c, "Code is not sent please try again")
		return
	}*/
	c.JSON(200, "Your code is sent to "+body.Email)

}

func logout(c *gin.Context) {
	token := c.GetHeader("token")
	fmt.Println(token)
}

func something(c *gin.Context) {
	token := c.GetHeader("Token")
	verifyJwt := helpers.VerifyJWT(token, c)
	fmt.Println(verifyJwt)

}
