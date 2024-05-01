package adapter

import (
	"docker-example/src/ports/in/drive/web/dto"
	"docker-example/src/ports/in/handler"

	"github.com/gin-gonic/gin"
)

func GinAdapter(fn handler.ContractHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := c.GetRawData()
		headers := map[string][]string{}
		params := map[string][]string{}

		for key, value := range c.Request.Header {
			headers[key] = value
		}

		for key, value := range c.Request.URL.Query() {
			params[key] = value
		}

		for _, value := range c.Params {
			params[value.Key] = []string{value.Value}
		}

		response, err := fn(params, params, headers, body)

		if err != nil {
			c.JSON(err.Status(), err.Body())
			return
		}
		resultResponse := response.(*dto.DefaultResponse)
		for key, value := range resultResponse.ResponseHeaders {
			c.Header(key, value)
		}
		c.JSON(resultResponse.ResponseCode, resultResponse.ResponseBody)
	}
}
