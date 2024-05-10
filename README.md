# qr-coder
[![codecov](https://codecov.io/gh/etilite/qr-coder/graph/badge.svg?token=A70ZRV50JV)](https://codecov.io/gh/etilite/qr-coder)

Microservice to generate QR-codes in png format.
Text content is encoded in UTF-8.

## Usage
### Build project
### Get docker Image
You can use pre-built Docker images to run `qr-coder`
```
docker run -it --rm -d -p 8080:8080 --name qr-coder \
  -e HTTP_ADDR=:8080 \
  etilite/qr-coder:latest
```

#### request 
`POST http://localhost:8080/generate`
```json
{
    "size": 256,
    "content": "https://github.com/etilite/qr-coder"
}
```
#### response
![response](https://github.com/etilite/qr-coder/assets/39223859/bd1c0946-905f-4244-9027-279e795bbdb3)

png image