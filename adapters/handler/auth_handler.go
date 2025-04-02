package handler

import (
	"gin-starter/core/dto"
	"gin-starter/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHandeler struct {
	sv ports.AuthService
}

func NewAuthHandler(service ports.AuthService) *authHandeler {
	return &authHandeler{sv: service}
}

var validate = validator.New()

const (
	CreateSuccessStatus   = http.StatusCreated
	BadRequestStatus      = http.StatusBadRequest
	InvalidRequestMessage = "inavalid JSON"
)

// @Summary     Sign Up
// @Description Register a new user and return JWT token
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       user body dto.SignUpRequest true "Sign Up Request"
// @Success     200 {object} dto.SignUpResponse
// @Failure     400 {object} errs.AppError
// @Failure     417 {object} errs.AppError
// @Router      /v1/auth/sign-up [post]
func (a *authHandeler) SignUp(c *gin.Context) {

	signUpRequest := new(dto.SignUpRequest)

	if err := c.ShouldBindJSON(signUpRequest); err != nil {
		c.JSON(BadRequestStatus, gin.H{
			"message": InvalidRequestMessage,
		})
		return
	}

	if err := validate.Struct(signUpRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		validateError := TranslateError(errors)
		HandlerError(c, validateError)
		return
	}

	resutl, err := a.sv.SignUp(signUpRequest)
	if err != nil {
		HandlerError(c, err)
		return
	}

	respone := dto.SignUpResponse{
		Message: "Sign Up successfully.!",
		Token:   resutl,
	}

	c.JSON(CreateSuccessStatus, respone)
}

// @Summary     Login
// @Description Authenticate a user and return JWT token
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       credentials body dto.LoginRequest true "Login Request"
// @Success     200 {object} dto.TokenResponse
// @Failure     400 {object} errs.AppError
// @Failure     417 {object} errs.AppError
// @Router      /v1/auth/login [post]
func (a *authHandeler) Login(c *gin.Context) {
	loginRequest := new(dto.LoginRequest)

	if err := c.ShouldBindJSON(loginRequest); err != nil {
		c.JSON(BadRequestStatus, gin.H{
			"message": InvalidRequestMessage,
		})
		return
	}
	if err := validate.Struct(loginRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		HandlerError(c, errors)
		return
	}

	result, err := a.sv.Login(loginRequest)
	if err != nil {
		HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
