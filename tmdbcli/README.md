# TMDB CLI

A command-line interface for The Movie Database (TMDB) API. To solve https://roadmap.sh/projects/tmdb-cli.

## Features

- Fetch top rated, now playing, popular, and upcoming movies from TMDB
- Pretty prints results in a table format
- Command-line flags for movie type selection
- Environment variable support for API key

## Environment Variables

- `TMDB_API_KEY` (required): Your TMDB API key. Set this in your `.env` file or environment.

## Setup

### Prerequisites

- Go 1.23.3 or later
- [TMDB API Key](https://developer.themoviedb.org/docs)
- Visual Studio Code (suggested)
- Go extension for VSCode (suggested)
- [Make](https://www.gnu.org/software/make/) for task automation
- [Task](https://taskfile.dev/) for task automation

### Environment Setup

1. Clone the repository
2. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
3. Add your TMDB API key to the `.env` file

## Supported Commands

- `movies --type top` - Get top rated movies
- `movies --type playing` - Get now playing movies
- `movies --type popular` - Get popular movies
- `movies --type upcoming` - Get upcoming movies

### Running the Application

Using Make:
```bash
make topcmd  # Run a command to get top rated movies
```

Using Task:
```bash
task topcmd  # Run a command to get top rated movies
```

Using Go directly:
```bash
go run main.go movies --type top  # Run a command to get top rated movies
```

### Building

```bash
go build .  # Creates tmdbcli binary
```

## Project Structure

- `cmd/` - CLI commands and flags
- `internal/tmdb/` - TMDB API client logic
- `main.go` - Application entry point
