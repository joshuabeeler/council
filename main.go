package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"fmt"

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

		// TODO: If there isn't a scheduled event, show to-be-determined info.

		loc, err := time.LoadLocation("America/Los_Angeles")
		if err != nil {
			log.Print(err)
		}

		startTime := time.Date(2020, 5, 4, 21, 0, 0, 0, loc)
		openMinutesBefore, _ := time.ParseDuration("-10m")
		openTime := startTime.Add(openMinutesBefore)
		endMinutesAfter, _ := time.ParseDuration("1.5h")
		endTime := startTime.Add(endMinutesAfter)

		dateStr := startTime.Format("Monday, January 2, 2006")
		log.Print(dateStr)

		openTimeClean := openTime.Format("304PM")

		// TODO: Phase out this service. Come up with a better method.
		timeZoneQuery := fmt.Sprintf(
			"https://time.is/%s_%d_%s_%d_in_Los_Angeles?Online_Council",
			openTimeClean, startTime.Day(), startTime.Format("Jan"), startTime.Year())

		zone, offset := startTime.Zone()
		offset /= 60 * 60
		log.Print(zone)
		log.Print(offset)
		zoneStr := fmt.Sprintf(
			"%s %d %s",
			zone, offset, startTime.Location().String())

		scheduledEvent := struct {
			Title string
			NextDate string
			TimeZone string
			OpenTime string
			StartTime string
			EndTime string
			ZoomLink string
			ZoomNumber string
			TimeZoneQuery string
		}{
			"Online Council",
			dateStr,
			zoneStr,
			openTime.Format("3:04 pm"),
			startTime.Format("3:04 pm"),
			endTime.Format("3:04 pm"),
			"TODO",
			"TODO",
			timeZoneQuery,
		}
		
		c.HTML(http.StatusOK, "index.gohtml", scheduledEvent)
	})

	router.Run(":" + port)
}
