# baseimage
FROM golang:1.18.3-alpine3.16

RUN mkdir /app

WORKDIR /app

# WORKDIR /app/go-modules
# Download necessary Go modules
# COPY ./src/go.mod ./
# COPY ./src/go.sum ./
# COPY /src ./
ADD /src ./
# ADD . .
RUN go mod download



# EXPOSE 8081
RUN ls

# CMD ["cd /src/cmd/web"]

# RUN go build -o /handlers internal/handlers/handlers.go
RUN go build -o /my-app  cmd/web/*.go
# ENTRYPOINT go build cmd/web/*.go && ./app
# RUN go run cmd/web/*.go
CMD [ "/my-app" ]