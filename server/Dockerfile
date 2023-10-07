FROM golang:1.21-alpine

WORKDIR /server

RUN go install github.com/cosmtrek/air@latest

# Copying files to work directory
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

EXPOSE 8080

# CMD ["./app"]
CMD ["air", "-c", ".air.toml"]
