package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		// Get the user's time zone from the query parameters
		timeZone := r.URL.Query().Get("tz")
		if timeZone == "" {
			timeZone = "UTC"
		}

		// Get the current time
		currentTime := time.Now().In(time.FixedZone(timeZone, 0))

		// Format the time and location as a string
		timeString := fmt.Sprintf("%s, %s, %s, %s",
			currentTime.Weekday(),
			currentTime.Format("15:04:05"),
			currentTime.Location(),
			currentTime.Format("-0700 MST"))

		// Write the string to the response writer
		w.Write([]byte(timeString))
	})

	http.ListenAndServe(":8080", nil)
}