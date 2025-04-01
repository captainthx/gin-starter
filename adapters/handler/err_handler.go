package handler

import (
	"errors"
	"fmt"
	errs "gin-starter/common/err"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandlerError(c *gin.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.JSON(e.Code, gin.H{
			"error": e.Message,
		})
	case error:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
	}
}

func TranslateError(errs validator.ValidationErrors) error {
	errorMsg := ""
	for _, e := range errs {
		field := e.Field()
		tag := e.Tag()
		param := e.Param()

		switch tag {
		case "required":
			errorMsg = fmt.Sprintf("%s is required", field)
		case "min":
			errorMsg = fmt.Sprintf("%s must be at least %s characters long", field, param)
		case "max":
			errorMsg = fmt.Sprintf("%s must be at most %s characters long", field, param)
		case "email":
			errorMsg = "Invalid email format"
		case "alphanum":
			errorMsg = fmt.Sprintf("%s must contain only letters and numbers", field)
		default:
			errorMsg = fmt.Sprintf("%s is invalid", field)
		}

		if errorMsg != "" {
			break
		}

	}
	return errors.New(errorMsg)
}
