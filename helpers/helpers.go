package helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// ValidateLen not really necessary but I like making function
func ValidateLen(validate string, length int) bool {
	if len(validate) < length {
		return true
	}
	return false
}

// AcceptMethod overlaps almost every method call so, it is better to make a function that we can use in every route.
func AcceptMethod(body interface{}, c *gin.Context) bool {
	data, err := c.GetRawData()
	if err != nil {
		NotAcceptableAbort(c, "Input format is wrong")
		return false
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		BadRequestAbort(c, "Can't match with the struct")
		return false
	}
	return true

}
