package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://763f4dfb120449d3a8a0f1c1e1719723@o398852.ingest.sentry.io/5705647",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("Go-Slip service is started")

	fmt.Println("Go-Slip Service is running...")

}
