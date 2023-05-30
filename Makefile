PROGRAM_NAME = app
REGISTRY = b0shka
API_IMAGE = walkom-backend
TAG = stable

.PHONY: build start run clean docker-build docker-run docker-push
.DEFAULT_GOAL := build

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/${PROGRAM_NAME} ./cmd/app/main.go

start: build
	.bin/app

run:
	go run ./cmd/app/main.go

clean:
	rm .bin/${PROGRAM_NAME}
	rmdir .bin

docker-build:
	docker build -f deploy/Dockerfile -t ${REGISTRY}/${API_IMAGE}:${TAG} .

docker-run:
	docker run -d -p 8000:8000 -e GIN_MODE=release --rm --name ${API_IMAGE} ${REGISTRY}/${API_IMAGE}:${TAG}

docker-push:
	docker push ${REGISTRY}/${API_IMAGE}:${TAG}