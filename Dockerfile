FROM golang:1.21

WORKDIR /usr/src/app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest


# Add files
ADD . .

RUN go mod download

RUN go build /usr/src/app/cmd/API/main.go

EXPOSE 8080

CMD go run /usr/src/app/cmd/API/main.go