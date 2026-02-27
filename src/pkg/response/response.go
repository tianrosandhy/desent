package response

import "github.com/labstack/echo/v4"

type Response struct {
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}

func ErrorResponse(message string, errData ...any) *Response {
	resp := Response{
		Message: message,
	}
	if len(errData) > 0 {
		resp.Error = errData[0]
	}

	return &resp
}

func (r *Response) Send(c echo.Context, httpCode int) error {
	return c.JSON(httpCode, r)
}
