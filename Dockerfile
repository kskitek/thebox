FROM golang:1.17.1-alpine3.13 AS builder

WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o web -ldflags="-s -w" cmd/web/main.go

FROM gcr.io/distroless/static
USER nonroot:nonroot

COPY --from=builder --chown=nonroot:nonroot /go/src/app/web /usr/bin/web

ENTRYPOINT ["web"]
