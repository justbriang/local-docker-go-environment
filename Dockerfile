FROM golang:alpine

RUN mkdir /app
WORKDIR /app

COPY /src ./

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o /my-app ./cmd/web" --command=/my-app
