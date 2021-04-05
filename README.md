# go-slip

Service for e-slip generator in golang. MIT license.

# Purpose

This project use to reduce slip paper for any POS or web application and let's customer or user choose to get an e-receipt via scan on-screen QR code or shortlink URL. So customer will redirect to service web and display or download slip file in PNG format for any purpose.

## How it's work.

Mobile Client or Caller service make a RESTful API call for NewSlip() function and service return a URL link for client to use for download or display a file

```mermaid
sequenceDiagram;
participant Caller
participant Service
  Caller->>Service: NewSlip()
  Note over Service: Create slip file on server
  Service->>Caller: Slip URL

```