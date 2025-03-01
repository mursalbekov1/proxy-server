# Proxy Server

Welcome to the Go Proxy Server, a simple HTTP proxy server written in Go. This server allows you to proxy HTTP requests to external services and retrieve their responses.

## Features

- Facilitates proxying of HTTP requests and returning responses.
- Docker support for containerization.

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/mursalbekov1/proxy-server.git
```

### Run and build the service:

```bash
make run
```

### Deployment link:

```bash
https://proxy-server-diro.onrender.com/
```

### Request
    
```json
    {
      "method": "GET",
      "url": "https://api.coindesk.com/v1/bpi/currentprice.json",
      "headers": {
        "Content-Type": "application/json"
      }
    }

```

### Response

```json
    {
      "id": 2,
      "status": 200,
      "headers": {
        "Accept-Ranges": "bytes",
        "Access-Control-Expose-Headers": "WWW-Authenticate,Server-Authorization",
        "Cache-Control": "max-age=30, must-revalidate, public",
        "Content-Length": "671",
        "Content-Type": "application/json; charset=utf-8",
        "Date": "Thu, 04 Jul 2024 02:03:42 GMT",
        "Vary": "origin",
        "Via": "1.1 a4fe306096165bb1e86e69365dc8fac2.cloudfront.net (CloudFront)",
        "X-Amz-Cf-Id": "wC0c3o9cG18Zfz7KHSXULi32G1gPPD5UNLU6pybYeJPg3h1HKWcZvg==",
        "X-Amz-Cf-Pop": "HIO50-C1",
        "X-Cache": "Miss from cloudfront"
      },
      "length": 671
    }
```