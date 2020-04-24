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

		// TODO: Make an event struct type.
		// TODO: If there isn't a scheduled event, show TBD template info.

		newYork, err := time.LoadLocation("America/New_York")
		if err != nil {
			log.Print(err)
		}

		startTime := time.Date(2020, 5, 4, 21, 0, 0, 0, newYork)
		openMinutesBefore, _ := time.ParseDuration("-10m")
		openTime := startTime.Add(openMinutesBefore)
		endMinutesAfter, _ := time.ParseDuration("1.5h")
		endTime := startTime.Add(endMinutesAfter)

		//dayOfWeek := "Wednesday" // TODO: Should be determined by year-month-day, above.

		dateStr := startTime.Format("Monday, January 2, 2006")
		log.Print(dateStr)

		//nextDate := fmt.Sprintf(
		//	"%s, %s %s, %s",
		//	date.Weekday().String(), month, day, year)

		// OPEN TIME
		//openTime := "5:50 pm"
		//openTimeClean := strings.Replace(openTime, ":", "", -1)
		//openTimeClean = strings.Replace(openTimeClean, " ", "", -1)
		//openTimeClean = strings.ToUpper(openTimeClean)
		openTimeClean := openTime.Format("304PM") // TODO: Can get rid of colon?

		// TODO: Phase out this service. Come up with a better method.
		// Maybe show PT, MT, CT, and ET in a table, w/ a link for other conversions?
		// Maybe give folks a drop-down (or search) to convert to their time zone?
		timeZoneQuery := fmt.Sprintf(
			"https://time.is/%s_%d_%d_%d_in_Los_Angeles?Online_Council",
			openTimeClean, startTime.Day(), startTime.Month(), startTime.Year())

		zone, offset := startTime.Zone()
		zoneStr := fmt.Sprintf(
			"%s %d %s",
			zone, offset, startTime.Location().String())

		scheduledEvent := struct {
			Title string			// TODO: Can probably remove.
			NextDate string		// TODO: Be a Date type.
			TimeZone string
			OpenTime string		// TODO: Should be StartTime - 10 minutes.
			StartTime string	// TODO: Should be a Time type.
			EndTime string		// TODO: Switch to duration. EndTime = StartTime + Duration
			ZoomLink string
			ZoomNumber string
			TimeZoneQuery string
		}{
			"Online Council",
			dateStr,
			zoneStr,
			openTime.Format("3:04 pm EDT"),		// 8:50 pm ET
			startTime.Format("3:04 pm EDT"),	// 9:00 pm ET
			endTime.Format("3:04 pm EDT"),	// 10:30 pm ET
			"TODO",
			"TODO",
			timeZoneQuery,
		}
		
		c.HTML(http.StatusOK, "index.gohtml", scheduledEvent)
	})

	router.Run(":" + port)
}
