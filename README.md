# github-tracker

Polls configured GitHub repositories for daily commit activity, caches reports, and delivers them via email through an async message queue.

Built to explore Go backend patterns: CQRS, hexagonal architecture, dependency injection via interfaces, and async pipelines.

## Stack

Go · Redis · RabbitMQ · SMTP

## How it works

GitHub API → HTTP handler → GenerateReport command ↓ Redis cache ↓ RabbitMQ (publish) ↓ Message consumer handler ↓ SendReport command ↓ SMTP email

## Architecture

Hexagonal layout with CQRS in the domain layer. Infrastructure concerns (HTTP, Redis, RabbitMQ, SMTP, file parsing) implement domain interfaces and are injected at startup in cmd/main.go.

cmd/main.go # wires dependencies, starts handlers

```
Internal/
    domain/
        interfaces.go -> JsonParser, RabbitMq, CacheRepo, Smtp, FileParser
        command/
            GenerateReport,
            SendReport,
        query/
            FetchAllReposFromConfFile
        formatter/
            Domain Logic Formatting Utils

    handler/
        http/
            FetchGithubData polls API, triggers report generation
        message/
            Send consumes queue, triggers email dispatch

    infrastructure/
        parser/
            JSON + file parsers
        rabbitmq/
            WorkQueue (Publish + Consume)
        email/
            SmtpClient

    repo/
        cache.go # Redis-backed CacheRepo

conf/ repos.toml # list of repositories to track
```

## Setup

```bash
git clone git@github.com:NewJhez01/github-tracker.git
cd github-tracker
cp .env.example .env          # fill in SMTP and connection details

docker run -d --rm --name redis -p 6379:6379 redis
docker run -d --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4-management
go run ./cmd/main.go

Environment

RABBIT_URL=amqp://guest:guest@localhost:5672/
REDIS_URL=localhost:6379
SMTP_FROM=
SMTP_HOST=
SMTP_PASSWORD=
SMTP_ADDR=

Remaining work

See open issues. Tests and cron scheduling are the main outstanding items.
```

```

```
