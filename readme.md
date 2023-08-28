# Generate QR code and export it as PDF

This project is an example how to generate QR code in PDF format using gin framework and Maroto module.

Under controller directory you will find 2 controller files that used to generate QR 
1. `generator.go` 
- `generator.go` is used to generate single page QR code within func `GenerateQR`
- Method: `GET`
- Endpoint: `/generate-qr`
- Parameter: nama (string) & nomor (string)

2. `multiplegenerator.go` 
- `multiplegenerator.go` is used to generate multiple pages of QR code within func `GenerateMultipleQR`
- Method: `POST` 
- Endpoint: `/multiple-qr`
- Request Body: json array

ex: 
```
[
    {
        "qrString": "TEST-123",
        "nomor": "TEST-123",
        "detail": "THIS IS MESSAGE TEST-123"
    },
    {
        "qrString": "TEST-456",
        "nomor": "TEST-456",
        "detail": "THIS IS MESSAGE TEST-456"
    },
    {
        "qrString": "TEST-789",
        "nomor": "TEST-789",
        "detail": "THIS IS MESSAGE TEST-789"
    }
]
```

## How To Run
Simply run command below on your terminal
```
go run main.go
```

the apps will run on your localhost with port 8081

## Dockerize this apps
Simply run command below on your terminal

```
docker build -t panic-qrcode:v1.0 .
```
