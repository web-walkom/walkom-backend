# WebAR-backend

The server part of the service "Walking in AR"

![GO][go-version]

---
## Installation

#### Requirements
* Golang 1.17
* Linux, Windows or macOS

#### Installing
```
git clone https://github.com/Danblack15/WebAR.git
cd WebAR
git checkout back
```

#### Configure
To work, you must create a `.env` file in the main directory of the project and specify such variables as:
```
MONDO_DB_URL - link to mongodb database
SALT - a combination of characters to generate a password hash
SECRET_KEY - key for generating authentication tokens
FRONTEND_URL - the link from which the request will come from the frontend
```

Also, in the `configs/config.yml` file, specify your mongodb login and the name of the database

---
## Usage
The port on which the service will be launched is specified in the file `configs/config.yml`

To start, run
```
go build cmd/app/main.go
./main
```

---
## Additionally
- The `api-walkom.service` file contains a setting for running backend on the server
- The `api-walkom.conf` file contains a setting for running backend using nginx on a subdomain `api.walkom.ru`
- The `walkom.service` file contains a setting for running frontend on the server
- The `walkom.conf` file contains a setting for running frontend using nginx on a domain `walkom.ru`


[go-version]: https://img.shields.io/static/v1?label=GO&message=v1.17&color=blue