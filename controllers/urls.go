package controllers

import (
	"github.com/Cguilliman/post-it-note/controllers/middlewares"
	"github.com/Cguilliman/post-it-note/controllers/notes"
	"github.com/Cguilliman/post-it-note/controllers/users"
	"github.com/gin-gonic/gin"
)

func InitRoutings(engine *gin.Engine) {
	// render := engine.Group("")
	v1 := engine.Group("/api/v1")
	v1.POST("/registration", users.UserRegisteration)
	v1.POST("/login", users.UserLogin)

	v1.Use(middlewares.AuthMiddleware(false))
	v1.GET("/retrieve", users.UserRetrieve)
	v1.PUT("/update", users.UserUpdate)
	v1.GET("/notes", notes.NotesList)
	v1.POST("/notes/create", notes.NoteCreate)
	v1.PUT("/note/:pk/update", notes.NoteUpdate)
	v1.DELETE("/note/:pk/delete", notes.NoteDelete)
}
