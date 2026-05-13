# Github Tracker

This git hub tracker will auto fetch all public activity by a given user. Generate an Email report and
asynchronously send it to the given person

## Description

This project is used to implement various backend strategies for me to improve and create a cool project at the same time.
Concepts used in this project are:
-Concurrent Http handline
-Concurrent Streaming of File
-Caching in Memory with Redis
-Async Message Handling
-Cron Job scheduling

## Getting Started

clone the project and then run the script setup.sh which will ask you for all the repos and write it into a toml file.
For more info on the functionality view the Architecture Diagramm below.

## Architecture

`┌─────────────────────────────────────────┐
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
│    [TOML]      [GitHub      [Cache]     │
│    (Repo File)  API]        (Report)    │
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
└─────────────────────────────────────────┘`

```

```
