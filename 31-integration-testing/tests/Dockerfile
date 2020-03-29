FROM golang:1.12.6

RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
CMD ["go", "test"]
