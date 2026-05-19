# syntax=docker/dockerfile:1

FROM golang:1.26.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-tracker ./cmd/main.go

EXPOSE 8080

CMD ["/docker-tracker"]
