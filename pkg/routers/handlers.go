package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
	FirstName string `json:"title"`
	LastName  string `json:"author"`
	Role      string `json:"description"`
}

func (h Handler) AddUser(ctx *gin.Context) {
	body := AddUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user User

	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Role = body.Role

	if result := h.DB.Create(&user); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &user)
}
