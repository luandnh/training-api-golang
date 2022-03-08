package v1

import (
	"strconv"
	"training-api-golang/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.IUser
}

func APIUser(r *gin.Engine, u service.IUser) {
	user := &UserHandler{
		userService: u,
	}
	group := r.Group("v1/user")
	{
		group.GET("", user.GetUsers)
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 50
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	code, result := h.userService.GetUsers(c, limitInt, offsetInt)
	c.JSON(code, result)
}
