package errorhandle

import (
	"net/http"
	"socmed/dto"
	"socmed/helper"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	}

	// Corrected the response assignment
	response := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	// Send the response
	c.JSON(statusCode, response)
}
