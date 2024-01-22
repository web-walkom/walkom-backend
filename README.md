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
MONDO_URI=mongodb://mongodb:27017
MONGO_DB_NAME=<db name>

EMAIL_SERVICE_NAME=<name complany>
EMAIL_SERVICE_ADDRESS=<email address>
EMAIL_SERVICE_PASSWORD=<email password>

SECRET_KEY=<random string>
```

---
## Build & Run
To start, run
```
make start
```


[go-version]: https://img.shields.io/static/v1?label=GO&message=v1.18&color=blue
