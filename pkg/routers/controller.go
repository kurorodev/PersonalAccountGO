package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &Handler{
		DB: db,
	}

	routes := router.Group("/users")
	routes.POST("/add", h.AddUser)
}
