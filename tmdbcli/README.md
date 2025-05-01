# TMDB CLI

A command-line interface for The Movie Database (TMDB) API. To solve https://roadmap.sh/projects/tmdb-cli.

## Setup

### Prerequisites

- Go 1.23.3 or later
- [TMDB API Key](https://developer.themoviedb.org/docs)
- Visual Studio Code (suggested)
- Go extension for VSCode (suggested)

### Development Tools

- [Task](https://taskfile.dev/) for task automation

### Environment Setup

1. Clone the repository
2. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
3. Add your TMDB API key to the `.env` file

### Running the Application

Using Task:
```bash
task topcmd  # Run a command to get top rated movies
```

Or using Go directly:
```bash
go run main.go movies --type top  # Run a command to get top rated movies
```

### Building

```bash
go build .  # Creates tmdbcli binary
```
