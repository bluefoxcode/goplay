package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// should be handled in its own routes world, but for now this will work as I learn.

	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/index.tmpl.html", nil)
	})

	router.Run(":" + port)

}
