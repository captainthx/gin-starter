package handler

import (
	"gin-starter/core/domain"
	"gin-starter/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
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
