# Environment
FROM golang:1.14 as build-env

RUN mkdir -p /opt/reg_service
WORKDIR /opt/reg_service
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /opt/service/reg_service

# Release
FROM alpine:latest
COPY --from=build-env /opt/service/reg_service /bin/reg_service
ENTRYPOINT ["/bin/reg_service"]
