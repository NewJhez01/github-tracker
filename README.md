# GitHub Tracker

Fetches public activity across configured repositories and delivers a weekly email digest.

## What it does

- Polls repository commits via GitHub API
- Generates a summary report
- Queues and sends it asynchronously via SMTP

## Concepts

- Concurrent HTTP workers
- File streaming (repo list from TOML)
- Redis caching
- Message queue + async consumers
- Cron scheduling

## Setup

```bash
git clone git@github.com:NewJhez01/github-tracker.git
./setup.sh  # writes repo list to config.toml
```

## Architecture

This project follows CQRS (Command Query Responsibility Segregation) with a hexagonal architecture mindset: domain logic sits at the center, handlers orchestrate, and infrastructure concerns (file I/O, HTTP, cache, queue, SMTP) are pushed to the edges.

```
┌─────────────────────────────────────────────────────────────┐
│                        Entry Point                          │
│                    cmd/tracker/main.go                      │
│              (wires dependencies, starts cron)              │
└─────────────────────────────┬───────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                         Handlers                            │
│  ┌─────────────────┐  ┌─────────────────────────────────┐   │
│  │  FetchHandler   │  │      MessageConsumerHandler     │   │
│  │                 │  │                                 │   │
│  │  Load repos ──→ │  │  ←── Read from queue            │   │
│  │  Fetch commits  │  │  Format report                  │   │
│  │  Generate report│  │  Send via SMTP                  │   │
│  │  Cache + Queue  │  │                                 │   │
│  └────────┬────────┘  └─────────────────────────────────┘   │
│           │                                                 │
│           │         ┌───────────────────┐                   │
│           └────────→│   ReportCreated   │←──────────────────┘
│                     │    (Domain Event) │                   │
│                     └───────────────────┘                   │
└─────────────────────────────┬───────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                          Domain                             │
│                                                             │
│   ┌──────────────┐        ┌──────────────────────────────┐  │
│   │    Query     │        │           Command            │  │
│   │              │        │                              │  │
│   │ GetRepos()   │        │  GenerateReport(repos)       │  │
│   │              │        │    → build commit summary    │  │
│   └──────────────┘        │    → return Report           │  │
│                           │                              │  │
│                           │  CacheReport(report)         │  │
│                           │  EnqueueReport(report)       │  │
│                           │                              │  │
│                           │  SendReport(report)          │  │
│                           │    → format email body       │  │
│                           │    → dispatch via SMTP       │  │
│                           └──────────────────────────────┘  │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │              internal/domain/model/                 │   │
│   │   Repo, Commit, Report, ReportCreatedEvent          │   │
│   └─────────────────────────────────────────────────────┘   │
└─────────────────────────────┬───────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Infrastructure                         │
│                                                             │
│  ┌────────────┐  ┌────────────┐  ┌──────────┐  ┌────────┐   │
│  │   File     │  │   GitHub   │  │  Redis   │  │ Queue  │   │
│  │  Reader    │  │   Client   │  │  Cache   │  │        │   │
│  │            │  │            │  │          │  │        │   │
│  │ TOML parse │  │ Commits API│  │ Store    │  │ Enqueue│   │
│  │            │  │ Rate limit │  │ Retrieve │  │ Dequeue│   │
│  └────────────┘  └────────────┘  └──────────┘  └────────┘   │
│                                                             │
│  ┌────────────┐  ┌──────────────────────────────────────┐   │
│  │   SMTP     │  │           internal/services/         │   │
│  │  Sender    │  │  (parsers, formatters, utilities —   │   │
│  │            │  │   pure functions, no I/O)            │   │
│  │ Send email │  │                                      │   │
│  └────────────┘  └──────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

## Package Layout Target

```
github-tracker/
├── cmd/
│   └── tracker/
│       └── main.go                 # entry point, wires deps
│
├── internal/
│   ├── handler/
│   │   ├── fetch.go                # FetchHandler
│   │   └── message_consumer.go     # MessageConsumerHandler
│   │
│   ├── domain/
│   │   ├── model/
│   │   │   ├── repo.go
│   │   │   ├── commit.go
│   │   │   └── report.go
│   │   │
│   │   ├── query/
│   │   │   └── repos.go            # GetRepos(repoReader) []Repo
│   │   │
│   │   └── command/
│   │       ├── generate_report.go
│   │       ├── cache_report.go
│   │       └── send_report.go
│   │
│   ├── infrastructure/
│   │   ├── file/
│   │   │   └── reader.go           # os.Open, TOML parse
│   │   ├── github/
│   │   │   └── client.go           # http.Client, API calls
│   │   ├── cache/
│   │   │   └── redis.go            # Redis client wrapper
│   │   ├── queue/
│   │   │   └── queue.go            # enqueue/dequeue
│   │   └── smtp/
│   │       └── sender.go           # email dispatch
│   │
│   └── services/
│       ├── parse_repos.go          # string splitting (pure)
│       ├── parse_commits.go        # JSON → domain model (pure)
│       └── format_email.go         # Report → HTML/text (pure)
│
├── conf/
│   └── repos.toml
│
└── setup.sh
```
