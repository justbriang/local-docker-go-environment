# baseimage
FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY *.go ./

# WORKDIR /app/go-modules
# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

# EXPOSE 8081

RUN go build -o /my-app
CMD [ "/my-app" ]