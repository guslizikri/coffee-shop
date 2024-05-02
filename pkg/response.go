package pkg

import (
	"coffee-shop/config"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code        int         `json:"-"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data,omitempty"`
	Meta        interface{} `json:"meta,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
	ctx.Abort()
}

func NewRes(code int, data *config.Result) *Response {
	var response = Response{
		Code:   code,
		Status: getStatus(code),
	}

	if response.Code >= 400 {
		response.Description = data.Data
	} else if data.Message != nil {
		response.Description = data.Message
	} else {
		response.Data = data.Data
	}

	if data.Meta != nil {
		response.Meta = data.Meta
	}

	return &response
}

func getStatus(code int) (desc string) {

	switch code {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 202:
		desc = "Accepted"
	case 304:
		desc = "Not Modified"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 403:
		desc = "Forbidden"
	case 404:
		desc = "Not Found"
	case 415:
		desc = "Unsupported Media Type"
	case 500:
		desc = "Internal Server Error"
	case 502:
		desc = "Bad Gateway"
	default:
		desc = "Status Code Undefined"
	}

	return

}
