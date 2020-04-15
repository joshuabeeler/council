package main

import (
	"log"
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
	//router.Static("/static", "static")
	router.Static("/www", "")

	//router.GET("/", func(c *gin.Context) {
	//	//c.HTML(http.StatusOK, "index.tmpl.html", nil)
	//	c.HTML(http.StatusOK, "www/index.html", nil)
	//})

	router.Run(":" + port)
}
