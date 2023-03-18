# walkom-backend

The server part of the service [walkom.ru](https://walkom.ru)

![Go][go-version]

---
## Installation

#### Prerequisites
- Go 1.18
- Docker
- Linux, Windows or macOS

Create `.env` file in root directory and add following values:
```
MONDO_URI = mongodb://mongodb:27017
MONGO_USER = admin
MONGO_PASSWORD = qwerty
FRONTEND_URL = http://localhost:3000
```

To install all the dependencies, run
```
go mod download
```

Also, in the `configs/config.yml` file, specify your mongodb database name

---
## Build & Run
The port on which the service will be launched is specified in the file `configs/config.yml`

To start, run
```
make start
```


[go-version]: https://img.shields.io/static/v1?label=GO&message=v1.18&color=blue
