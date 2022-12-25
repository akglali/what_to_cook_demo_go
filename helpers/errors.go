package helpers

import "github.com/gin-gonic/gin"

type errorStruct struct {
	Error string
}

func BadRequestAbort(c *gin.Context, errorMessage string) {
	//400 The server cannot or will not process the request due to something that is perceived to be a client error
	//(e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
	c.AbortWithStatusJSON(400, errorStruct{Error: errorMessage})
}

func NotAcceptableAbort(c *gin.Context, errorMessage string) {
	//406 This response is sent when the web server, after performing server-driven content negotiation,
	//doesn't find any content that conforms to the criteria given by the user agent.
	c.AbortWithStatusJSON(406, errorStruct{Error: errorMessage})
}
