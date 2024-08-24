# qr-coder
[![docker pulls](https://img.shields.io/docker/pulls/etilite/qr-coder)](https://hub.docker.com/r/etilite/qr-coder)
[![docker push](https://github.com/etilite/qr-coder/actions/workflows/docker.yml/badge.svg)](https://github.com/etilite/qr-coder/actions/workflows/docker.yml)
[![go build](https://github.com/etilite/qr-coder/actions/workflows/go.yml/badge.svg)](https://github.com/etilite/qr-coder/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/etilite/qr-coder/graph/badge.svg?token=A70ZRV50JV)](https://codecov.io/gh/etilite/qr-coder)

Lightweight microservice written in Go to generate QR-codes in `png` format.

By default, `qr-coder` adds to content `UTF-8 BOM` prefix for better unicode-compatibility with scanners.

## Usage
### Quick Start with Docker

```sh
docker run --rm -p 8080:8080 -e HTTP_ADDR=:8080 etilite/qr-coder:latest
```

This will start the service and expose its API on port 8080.

### API

#### Request
- `POST /generate`
```json
{
  "size": 256,
  "content": "https://github.com/etilite/qr-coder"
}
```
- `size` in pixels of image side, `int`
- `content` is a string to encode, `string`

**Request Example:**

Using `cURL`, you can make a request like this:

```sh
curl --location 'localhost:8080/generate' \
--header 'Content-Type: application/json' \
--data '{
  "size": 256,
  "content": "https://github.com/etilite/qr-coder"
}' -o img.png
```

#### Response
The response will be a binary PNG file with the encoded content.

![response](https://github.com/etilite/qr-coder/assets/39223859/bd1c0946-905f-4244-9027-279e795bbdb3)

### Build from source

```sh
git clone https://github.com/etilite/qr-coder.git
cd qr-coder
make run
```
This will build and run app at `http://localhost:8080`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

If you'd like to contribute to the project, please open an issue or submit a pull request on GitHub.