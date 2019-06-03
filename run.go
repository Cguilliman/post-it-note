package main

import (
	"github.com/Cguilliman/post-it-note/controllers"
	"github.com/Cguilliman/post-it-note/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := models.InitDatabase()
	defer db.Close() // close connection after server stopped
	engine := gin.Default()
	// engine.LoadHTMLGlob("templates/*")
	controllers.InitRoutings(engine)
	// testDbWorking(db)
	engine.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:9000
}
