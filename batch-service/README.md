# Batch Service using gobatch

> **Status: parked — proof of concept.**
> This module was a learning exercise to explore the Spring Batch
> (reader → processor → writer, chunked steps, partitions, job metadata)
> model in Go. It works end-to-end but is **not maintained**: the underlying
> library [`chararch/gobatch`](https://github.com/chararch/gobatch) is itself
> unmaintained (last release v1.0.2, ~2023). Kept as a reference for the
> pattern, not as a basis for further work. For production-grade batch work in
> Go, prefer composing `golang.org/x/sync/errgroup` + channels with a bounded
> worker pool (`ants`/`pond`) over adopting a dead framework.

## What it does

`job1` reads a CSV of ~1,000 persons, maps each record to a `PersonV1Entity`,
and persists the batch into the MySQL `persons` table. Verified: 1000/1000 rows
inserted, job status `COMPLETED`.

## Tooling

- Go 1.26
- [GoBatch](https://github.com/chararch/gobatch) v1.0.2
- MySQL 8.0

## Running locally

Configuration is read from environment variables (`DB_USERNAME`, `DB_PASSWORD`,
`DB_HOST`, `DB_PORT`, `DB_NAME`), falling back to local defaults — see
`db/config.go`.

```sh
cd _devenvironment

# The compose plugin needs the image pulled by podman first (avoids the
# osxkeychain credential helper), then --env-file interpolates ${DB_NAME}.
podman pull docker.io/library/mysql:8.0.32
podman compose --env-file compose.env up -d

# The schemas are NOT auto-loaded by compose — load them manually:
podman exec -i devenvironment-db-1 mysql -uroot gobatchservicedb < db/00_schema_mysql_gobatch.sql
podman exec -i devenvironment-db-1 mysql -uroot gobatchservicedb < db/01_schema_mysql_application.sql

cd ..
go run .
```

> Note: `main.go` runs the job with `Partitions(100)`. The connection pools are
> bounded (`SetMaxOpenConns`) so the concurrent partitions queue instead of
> exhausting MySQL's `max_connections` (`Error 1040: Too many connections`).
