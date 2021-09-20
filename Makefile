CGO_ENABLED=0
DOCKER_BUILDKIT=1
REPOSITORY=eu.gcr.io/skitek/thebox
VERSION=0.1
IMAGE=$(REPOSITORY):$(VERSION)

build: build-web build-box

build-web:
	go build -o web cmd/web/main.go

build-box:
	go build GOARCH=arm64 GOOS=linux -o box cmd/box/main.go

docker-build:
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)

env:
	cat .env.template | grep -v '#' > .env
