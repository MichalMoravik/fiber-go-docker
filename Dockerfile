FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ /app/src

RUN go build -o server ./src
EXPOSE 8080

CMD ["/app/server"]
