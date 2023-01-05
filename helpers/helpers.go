package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AcceptMethod overlaps almost every method call so, it is better to make a function that we can use in every route.
func AcceptMethod(body interface{}, c *gin.Context) bool {
	data, err := c.GetRawData()
	if err != nil {
		NotAcceptableAbort(c, "Input format is wrong")
		return false
	}

	if err := json.Unmarshal(data, &body); err != nil {
		fmt.Println(err)
		BadRequestAbort(c, "Can't match with the struct")
		return false
	}

	if err := validator.New().Struct(body); err != nil {
		BadRequestAbort(c, err.Error())
		return false
	}

	return true

}
