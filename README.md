# Spotify Toolbox

A CLI tool for managing Spotify podcasts in bulk, built with Go and Cobra framework.

## Features

- List all podcasts in your Spotify library
- Delete podcasts in bulk with filters
- Dry-run mode to preview changes before applying

## Prerequisites

- Spotify Developer Account
- Spotify Client ID and Client Secret

## Setup

### 1. Get Spotify API Credentials

1. Go to [Spotify Developer Dashboard](https://developer.spotify.com/dashboard)
2. Create a new application
3. Copy your Client ID and Client Secret

### 2. Configure Environment Variables

Copy the example environment file and fill in your credentials:

```bash
cp .env.example .env
```

Edit `.env` and add your Spotify credentials:

```env
SPOTIFY_CLIENT_ID=your_actual_client_id
SPOTIFY_CLIENT_SECRET=your_actual_client_secret
```

Alternatively, you can create a config file at `~/.spotify-toolbox.yaml`:

```yaml
client_id: your_actual_client_id
client_secret: your_actual_client_secret
```

### 3. Build the Application

```bash
make build
```

Or manually:

```bash
go build -o bin/spotify-toolbox ./cmd/spotify-toolbox
```

## Usage

### List all podcasts

```bash
./bin/spotify-toolbox list
```

With options:

```bash
# Output as JSON
./bin/spotify-toolbox list --format json

# Limit number of results
./bin/spotify-toolbox list --limit 100
```

### Delete podcasts

```bash
# Delete all podcasts (with confirmation)
./bin/spotify-toolbox delete --all

# Delete with filter
./bin/spotify-toolbox delete --filter "podcast-name-pattern"

# Dry run to see what would be deleted
./bin/spotify-toolbox delete --all --dry-run
```

### Using flags for credentials

Instead of environment variables or config file, you can pass credentials as flags:

```bash
./bin/spotify-toolbox --client-id YOUR_ID --client-secret YOUR_SECRET list
```

## Development

### Install dependencies

```bash
go mod tidy
```

### Run tests

```bash
# Run all tests
make test

# Run tests with race detector
make test-race

# Run tests with coverage
make test-coverage
```

### Code Quality

```bash
# Format code
make fmt

# Run static analysis
make vet

# Run linter (requires golangci-lint)
make lint

# Run all checks
make check
```

### Build

```bash
make build
```

### Clean build artifacts

```bash
make clean
```

## Project Structure

```text
spotify-toolbox/
├── cmd/
│   └── spotify-toolbox/    # Application entry point
│       └── main.go
├── internal/
│   ├── cmd/                # Cobra commands
│   │   ├── root.go
│   │   ├── list.go
│   │   └── delete.go
│   ├── spotify/            # Spotify API client
│   └── config/             # Configuration management
├── pkg/
│   └── logger/             # Logging utilities
├── .env.example            # Environment variables template
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## License

MIT

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.
