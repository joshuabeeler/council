package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	// From: https://github.com/gin-gonic/gin#installation
	router.Use(gin.Recovery())

	//router.LoadHTMLGlob("templates/*.tmpl.html")
	router.LoadHTMLGlob("templates/main.tmpl.html")

	//router.Static("/static", "static")
	router.StaticFS("/static-www", http.Dir("static-www"))

	router.GET("/", func(c *gin.Context) {
		log.Print("Requested GET /")
		//c.HTML(http.StatusOK, "index.tmpl.html", nil)
		c.HTML(http.StatusOK, "templates/main.tmpl.html", struct {
			Title string
		}{
			"Hello, World",
		})
	})

	router.Run(":" + port)
}
