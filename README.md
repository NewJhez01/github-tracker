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

## Architecture Diagram

```
┌─────────────────────────────────────────┐
│              Cron Trigger               │
│         (systemd timer / scheduler)     │
└─────────────────┬───────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────┐
│           Go App (Main)                 │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐  │
│  │ Handler │→ │ Handler │→ │ Handler │  │
│  │  Load   │  │  Fetch  │  │ Generate│  │
│  │  Config │  │  GitHub │  │  Report │  │
│  └────┬────┘  └────┬────┘  └────┬────┘  │
│       │            │            │       │
│       ▼            ▼            ▼       │
│    [TOML]     [GitHub API]  [Cache]     │
│    (Repo File)              (Report)    │
│                              │          │
│                              ▼          │
│                         [Queue]         │
│                              │          │
│                              ▼          │
│  ┌──────────────────────────────────┐   │
│  │        Consumer Handler          │   │
│  │    Read Cache → Format Email     │   │
│  │         → Send via SMTP          │   │
│  └──────────────────────────────────┘   │
└─────────────────────────────────────────┘
```
