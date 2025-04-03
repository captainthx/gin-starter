package handler

import (
	"gin-starter/core/domain"
	"gin-starter/core/dto"
	"gin-starter/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	sv ports.UserService
}

func NewUserHandler(service ports.UserService) *userHandler {
	return &userHandler{sv: service}
}

// @Summary     User info
// @Description Get user info From JWT token
// @Tags        User
// @Accept      json
// @Produce     json
// @Success     200 {object} dto.UserInfo
// @Failure     404 {object} errs.AppError
// @Router      /v1/user/me [get]
func (u *userHandler) GetUser(c *gin.Context) {
	user := c.MustGet("user").(*domain.User)

	result, err := u.sv.GetAccount(user.ID)
	if err != nil {
		HandlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u *userHandler) UpdataAvatar(c *gin.Context) {
	updAvatarRequest := new(dto.UpdateAvartarRequest)

	if err := c.ShouldBindJSON(&updAvatarRequest); err != nil {
		c.JSON(BadRequestStatus, gin.H{
			"message": InvalidRequestMessage,
		})
		HandlerError(c, err)
		return
	}

	if err := Validate.Struct(updAvatarRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		validateError := TranslateError(errors)
		HandlerError(c, validateError)
		return
	}
	user := c.MustGet("user").(*domain.User)

	err := u.sv.UpdateAvatar(updAvatarRequest.AvatarUrl, user.ID)
	if err != nil {
		HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update avatart successfully",
	})
}
