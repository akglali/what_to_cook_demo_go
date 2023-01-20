package user

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
	"net/smtp"
	"os"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"
const sqlDateFormat = "2006-01-02T15:04:05Z"

// hash the password for users safety
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

// it checks the hashed password
func checkPassword(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

// it generates the user token
func tokenGenerator() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	//time is added to make token more unique
	currentTime := time.Now().Format("Mon Jan _2 15:04:05 2006")
	b = append(b, []byte(currentTime)...)

	return hex.EncodeToString(b)
}

// i use it to create a random code
const otpChars = "1234567890"

// GenerateOTP 6 digits code is generated!
func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

// send email you need to get mail address and password by creating .env file
func sendEmail(subject, body, email string) error {

	from := os.Getenv("MAIL")
	password := os.Getenv("PASSWD")
	host := "smtp.gmail.com"
	port := 587

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", from, password, host)
	to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, from, to, msg)
	return err
}

// check if the code is expired
func timeDiffSec(createdAt string) bool {
	currentTime := time.Now()

	// we need to convert the time location. To make the right calculation
	loc := currentTime.Location()
	date, err := time.ParseInLocation(sqlDateFormat, createdAt, loc)
	if err != nil {
		return false
	}
	diff := currentTime.Sub(date).Seconds()
	if diff > 300 {
		return false
	}
	return true

}

// google sign up method
func validateGoogleAccount(idToken string) {
	payload, err := idtoken.Validate(context.Background(), idToken, os.Getenv("GOOGLE-CLIENT-ID"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(payload.Claims)
}
