package main

import (
	"log"
	"net/http"
	"os"

	"fmt"
	"strings"

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

		// TODO: Make an event struct type.
		// TODO: If there isn't a scheduled event, show TBD template info.

		// DATE
		dayOfWeek := "Wednesday"
		month := "Apr"
		day := "4"
		year := "2020"

		nextDate := fmt.Sprintf("%s, %s %s, %s", dayOfWeek, month, day, year)
		
		// OPEN TIME
		openTime := "5:50 pm"
		openTimeClean := strings.Replace(openTime, ":", "", -1)
		openTimeClean = strings.Replace(openTimeClean, " ", "", -1)
		openTimeClean = strings.ToUpper(openTimeClean)

		timeZoneQuery := fmt.Sprintf("https://time.is/%s_%s_%s_%s_in_Los_Angeles?Online_Council", openTimeClean, day, month, year)

		c.HTML(http.StatusOK, "index.gohtml", struct {
			Title string				// TODO: Can probably remove.
			NextDate string			// TODO: Be a Date type.
			OpenTimePT string		// TODO: Be StartTimePT - 10 minutes.
			StartTimePT string	// TODO: Be a Time type.
			EndTimePT string		// TODO: Switch to duration. EndTimePT = StartTimePT + Duration
			ZoomLink string
			ZoomNumber string
			TimeZoneQuery string
		}{
			"Online Council",
			nextDate,
			openTime, // 8:50 pm ET
			"6:00 pm", // 9:00 pm ET
			"7:30 pm", // 10:30 pm ET
			"https://zoom.us/j/93622550259?pwd=ZlpiTmRFUHA4V0tHMHB1cEFubmIwZz09",
			"936-2255-0259",
			timeZoneQuery,
		})
	})

	router.Run(":" + port)
}
