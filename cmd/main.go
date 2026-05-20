package main

import (
	"log"
	"os"

	"NewJhez01/github-tracker/internal/handler/http"
	"NewJhez01/github-tracker/internal/handler/message"
	"NewJhez01/github-tracker/internal/infrastructure/email"
	"NewJhez01/github-tracker/internal/infrastructure/parser"
	"NewJhez01/github-tracker/internal/infrastructure/rabbitmq"
	"NewJhez01/github-tracker/internal/repo"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load env")
	}

	// connections
	rabbitConn, err := amqp091.Dial(os.Getenv("RABBIT_URL"))
	if err != nil {
		log.Fatalf("connection failed", err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	// parsers
	fParser := parser.NewFileParser()
	githubParser := parser.NewGithubParser()

	// clients
	rabbitmq, err := rabbitmq.NewPublisher(rabbitConn)
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq: %s", err.Error())
	}
	cr := repo.NewCacheRepo(redisClient)
	smtp := email.NewSmtpClient(
		os.Getenv("SMTP_FROM"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_ADDR"),
	)

	// currently for testing purposes until the cron job is active
	http.FetchGithubData(githubParser, rabbitmq, fParser, &cr)

	// open endless connection for message handler
	if len(os.Args) > 1 && os.Args[1] == "consume" {
		go message.Send(rabbitmq, &cr, smtp)
		select {}
	}
}
