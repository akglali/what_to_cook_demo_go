package user

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/smtp"
	"os"
	"time"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func checkPassword(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

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

const otpChars = "123456789"

// 6 digits code is generated!
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
