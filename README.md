# go-slip

Go Package Library for e-Slip generator. MIT license.
This package has dependency required WkHTML.

## Purpose

This project use to reduce slip paper for any POS or web application and let's customer or user choose to get an e-receipt via scan on-screen QR code or shortlink URL. So customer will redirect to service web and display or download slip file in PNG format for any purpose.

## How it's work

Mobile Client or Caller service make a RESTful API call for NewSlip() function and service return a Shorten-URL link for client to use for download or display a file

```mermaid
sequenceDiagram;
  participant Caller
  participant Service
  participant GCS
  Caller->>Service: Create Slip (json)
  Note over Service: Create slip file image
  Service->>GCS: Store image (png)
  GCS->>Service: Success
  Service->>Caller: Slip URL
```

## How to test

1. Create Google Cloud Storage - new Bucket such as "paybox_slip" as you see in "/slip/repository/gcs_test.go"

1. Create your Google Cloud - Service Account and place in "$HOME/secret/" and rename as paybox_slip_key.json 

1. then just run...

```bash
go test -v ./...
```

or to test just some package.

```bash
go test -v ./repository
```
