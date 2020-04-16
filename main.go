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

	router.LoadHTMLGlob("templates/*.gohtml")

	router.StaticFS("/static", http.Dir("static"))

	router.GET("/", func(c *gin.Context) {
		log.Print("Requested GET /")
		c.HTML(http.StatusOK, "index.gohtml", struct {
			Title string				// TODO: Can probably remove.
			NextDate string			// TODO: Be a Date type.
			OpenTimePT string		// TODO: Be StartTimePT - 10 minutes.
			StartTimePT string	// TODO: Be a Time type.
			EndTimePT string		// TODO: Switch to duration. EndTimePT = StartTimePT + Duration
			ZoomLink string
			ZoomNumber string
		}{
			"Online Council",
			"Wednesday, April 22, 2020",
			"5:50 pm", // 8:50 pm ET
			"6:00 pm", // 9:00 pm ET
			"7:30 pm", // 10:30 pm ET
			"https://zoom.us/j/93622550259?pwd=ZlpiTmRFUHA4V0tHMHB1cEFubmIwZz09",
			"936-2255-0259",
		})
	})

	router.Run(":" + port)
}
