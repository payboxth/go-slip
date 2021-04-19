package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/payboxth/go-slip"
	slipendpoint "github.com/payboxth/go-slip/endpoint"
	sliphandler "github.com/payboxth/go-slip/handler"
	sliprepository "github.com/payboxth/go-slip/repository"
	slipservice "github.com/payboxth/go-slip/service"
)

var (
	bucketName     string = "paybox_slip"
	credentialFile string = "/secret/paybox_slip_key.json"
	folderName     string = "image"
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

	database, err := sliprepository.NewBolt("slip.db")
	if err != nil {
		log.Fatal(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		sentry.CaptureMessage("Error from os.UserHomeDir():" + err.Error())
	}
	credentialFile = homeDir + credentialFile
	// Create new Google Cloud Storage(GCS) instant with Service Account Key in srcret folder.
	storage, err := sliprepository.NewGCS(bucketName, credentialFile)
	if err != nil {
		sentry.CaptureMessage(err.Error())
	}
	// Create new Slip Service and Endpoint
	service := slipservice.New(database, storage)
	endpoint := slipendpoint.New(service)

	mux := http.NewServeMux()
	mux.Handle("/", sliphandler.New(service))
	mux.Handle("/slip/", http.StripPrefix("/slip", slip.NewHTTPTransport(endpoint)))

	http.ListenAndServe(":8080", mux)

	sentry.CaptureMessage("Go-Slip service started.")
	fmt.Println("Go-Slip Service is running...")
}
