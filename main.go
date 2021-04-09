package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/payboxth/go-slip/slip"
	slipendpoint "github.com/payboxth/go-slip/slip/endpoint"
	sliphandler "github.com/payboxth/go-slip/slip/handler"
	sliprepository "github.com/payboxth/go-slip/slip/repository"
	slipservice "github.com/payboxth/go-slip/slip/service"
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

	database := sliprepository.NewBolt()
	storage := sliprepository.NewGCS()
	service := slipservice.New(database, storage)
	endpoint := slipendpoint.New(service)

	mux := http.NewServeMux()
	mux.Handle("/", sliphandler.New(service))
	mux.Handle("/slip/", http.StripPrefix("/slip", slip.NewHTTPTransport(endpoint)))

	http.ListenAndServe(":8080", mux)

	sentry.CaptureMessage("Go-Slip service is started")
	fmt.Println("Go-Slip Service is running...")

}
