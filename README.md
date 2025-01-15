# Mercury - Monolithic Server for Eventr in Go

## Setup Steps
This repository requires a Go installation. You would also require `make` 
available in your system.

```bash
git clone https://github.com/TeamEventr/mercury
cd mercury
docker-compose up -d  # Brings up Postgres, Redis and ElastiSearch
make run
```
For running the server in development mode with hot-reload, use 'air'. The 
configuration is already available in `.air.toml`

```bash
go install github.com/air-verse/air@latest
air
```
