package handler

import (
	"net/http"
	"technikom/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUser(c *gin.Context) {

	users, err := h.userService.GetUser()

	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
