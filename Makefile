CGO_ENABLED=0
DOCKER_BUILDKIT=1
REPOSITORY=eu.gcr.io/skitek/thebox
VERSION=0.1
IMAGE=$(REPOSITORY):$(VERSION)
PI_ADDR=192.168.0.111

build: build-web build-box

build-web:
	go build -o web cmd/web/main.go

build-box:
	GOARM=6 GOARCH=arm GOOS=linux go build -o box cmd/box/main.go
	upx -6 box

docker-build:
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)

upload-box:
	scp box pi@$(PI_ADDR):/home/pi

env:
	cat .env.template | grep -v '#' > .env
