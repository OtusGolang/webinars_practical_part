# Environment
FROM golang:1.14 as build-env

RUN mkdir -p /opt/notify_service
WORKDIR /opt/notify_service
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /opt/service/notify_service

# Release
FROM alpine:latest

COPY --from=build-env /opt/service/notify_service /bin/notify_service
ENTRYPOINT ["/bin/notify_service"]
