package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutings(engine *gin.Engine) {
	// render := engine.Group("")
	v1 := engine.Group("/api/v1")
	v1.POST("/registration", UserRegisteration)
	v1.POST("/login", UserLogin)

	v1.Use(AuthMiddleware(false))
	v1.GET("/retrieve", UserRetrieve)
	v1.PUT("/update", UserUpdate)
	v1.GET("/notes", NotesList)
	v1.PUT("/note/:pk/update", NoteUpdate)
	v1.DELETE("/note/:pk/delete", NoteDelete)
}
