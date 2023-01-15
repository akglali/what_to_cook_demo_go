package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"what_to_cook_demo_go/database"
	"what_to_cook_demo_go/user"
)

type signUpDTO struct {
	Username string `db:"username" form:"username" json:"username" validate:"required,min=6,max=18"`
	Password string `db:"password" form:"password" json:"password" validate:"required,min=6,max=18"`
	Email    string `db:"email" form:"email" json:"email" validate:"required"`
	Code     string `db:"code" form:"code" json:"code" validate:"required"`
}

type verifyUserDTO struct {
	Email string `db:"email" form:"email" json:"email" validate:"required"`
}

func TestDbConnection(t *testing.T) {
	database.ConnectDatabase()
}

var router *gin.Engine

func TestGinConnection(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
}

func TestVerifyUnit(t *testing.T) {
	router.POST("/Verify", user.VerifyAccount)
	emailAddress := "example@hotmail.com"
	email := verifyUserDTO{
		Email: emailAddress,
	}
	jsonValue, _ := json.Marshal(email)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/Verify", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	// Assertions
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"Your code is sent to "+emailAddress+"\"", w.Body.String())
}

func TestSignUp(t *testing.T) {
	router.POST("/", user.SignUp)
	signUp := signUpDTO{
		Email:    "example@hotmail.com",
		Username: "example",
		Password: "HelloWorld",
		Code:     "123566",
	}
	jsonValue, _ := json.Marshal(signUp)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	assert.Equal(t, 406, w.Code)

}
