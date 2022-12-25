package helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var (
	otpChars = "1234567890"
)

func ValidateLen(validate string, length int) bool {
	if len(validate) < length {
		return true
	}
	return false
}

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
