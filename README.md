# SSH Manager (sshmama)

[![Go Version](https://img.shields.io/github/go-mod/go-version/imrancluster/sshmama)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/imrancluster/sshmama)](https://goreportcard.com/report/github.com/imrancluster/sshmama)

A powerful command-line SSH connection manager that allows you to store, organize, and quickly connect to SSH hosts with encrypted private key support.

## ✨ Features

- 🔐 **Secure Storage**: Encrypted private key management using age encryption
- 🏷️ **Tagged Organization**: Organize connections with custom tags
- 🔍 **Smart Search**: Find connections quickly with fuzzy search
- 📱 **Interactive TUI**: Beautiful terminal interface for connection selection
- 💾 **Persistent Storage**: BoltDB-based storage for reliable data persistence
- 📤 **Import/Export**: Backup and restore your SSH configurations
- 🚀 **Fast Connections**: One-command SSH connections to any stored host

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/imrancluster/sshmama.git
cd sshmama

# Build the binary
make build
# or manually:
go build -o bin/sshmama ./cmd/sshmama

# Add to your PATH (optional)
sudo cp bin/sshmama /usr/local/bin/
```

### First Steps

```bash
# Add your first SSH connection
./bin/sshmama add --name prod-server --host 203.0.113.10 --user ubuntu --port 22

# Attach an encrypted private key
./bin/sshmama attach --name prod-server --file ~/.ssh/id_ed25519

# Connect to the server
./bin/sshmama connect prod-server

# Or use interactive selection
./bin/sshmama connect
```

## 📚 Commands Reference

### Core Commands

#### `sshmama add` - Add New SSH Connection

```bash
# Basic connection
sshmama add --name webserver --host 192.168.1.100 --user admin

# With custom port
sshmama add --name db-server --host 10.0.0.50 --user postgres --port 2222

# With tags and notes
sshmama add --name staging --host staging.example.com --user deploy --tags staging,web --notes "Staging environment for testing"
```

**Options:**
- `--name`: Connection name (required)
- `--host`: Hostname or IP address (required)
- `--user`: Username (required)
- `--port`: SSH port (default: 22)
- `--tags`: Comma-separated tags
- `--notes`: Additional notes

#### `sshmama list` - List All Connections

```bash
# List all connections
sshmama list

# List with details
sshmama list --verbose

# Filter by tags
sshmama list --tags production,web
```

#### `sshmama connect` - Connect to SSH Host

```bash
# Connect by name
sshmama connect webserver

# Interactive selection (when no name provided)
sshmama connect

# Dry run (show command without executing)
sshmama connect webserver --dry-run
```

#### `sshmama search` - Search Connections

```bash
# Search by name
sshmama search webserver

# Search by host
sshmama search 192.168.1

# Search by tags
sshmama search --tags production
```

### Key Management

#### `sshmama attach` - Attach Encrypted Private Key

```bash
# Attach private key file
sshmama attach --name webserver --file ~/.ssh/id_rsa

# The key will be encrypted and stored securely
# You'll be prompted for a passphrase
```

#### `sshmama edit` - Edit Connection Details

```bash
# Edit connection details
sshmama edit webserver

# Interactive editing of all fields
```

### Data Management

#### `sshmama export` - Export Connections

```bash
# Export to JSON
sshmama export --format json --output connections.json

# Export to CSV
sshmama export --format csv --output connections.csv
```

#### `sshmama import` - Import Connections

```bash
# Import from JSON file
sshmama import --file connections.json

# Import with conflict resolution
sshmama import --file connections.json --overwrite
```

#### `sshmama rm` - Remove Connections

```bash
# Remove by name
sshmama rm webserver

# Remove with confirmation
sshmama rm --confirm webserver
```

### Utility Commands

#### `sshmama completion` - Shell Completion

```bash
# Generate bash completion
sshmama completion bash > ~/.local/share/bash-completion/completions/sshmama

# Generate zsh completion
sshmama completion zsh > ~/.zsh/completions/_sshmama

# Generate fish completion
sshmama completion fish > ~/.config/fish/completions/sshmama.fish
```

## 🔧 Configuration

### Data Directory

By default, sshmama stores data in your OS configuration directory:

- **Linux/macOS**: `~/.config/sshmama/`
- **Windows**: `%APPDATA%\sshmama\`

You can specify a custom directory:

```bash
sshmama --data-dir /custom/path list
```

### Environment Variables

```bash
# Custom data directory
export sshmama_DATA_DIR="/custom/path"

# Debug mode
export sshmama_DEBUG="true"
```

## 🏗️ Architecture

```
sshmama/
├── cmd/sshmama/          # Main application entry point
├── internal/
│   ├── app/             # Application core and state management
│   ├── cli/             # Command-line interface implementations
│   ├── crypto/          # Age encryption for private keys
│   ├── db/              # BoltDB storage layer
│   ├── model/           # Data structures
│   ├── search/          # Search functionality
│   ├── ssh/             # SSH connection handling
│   └── util/            # Utility functions
└── pkg/version/         # Version information
```

## 🔐 Security

- **Private Key Encryption**: All private keys are encrypted using age encryption
- **Secure Storage**: Data is stored in a local BoltDB with proper file permissions
- **No Cloud Storage**: All data remains on your local machine
- **Passphrase Protection**: Each encrypted key requires a unique passphrase

## 📦 Dependencies

- **Go 1.24+**: Modern Go features and performance
- **BoltDB**: Embedded key-value store for data persistence
- **age**: Modern encryption for private key security
- **Cobra**: Powerful CLI framework
- **promptui**: Beautiful terminal user interface

## 🚀 Development

### Prerequisites

- Go 1.24 or later
- Git

### Building

```bash
# Build binary
make build

# Build with version info
make build-version

# Clean build artifacts
make clean

# Run tests
make test

# Install dependencies
make deps
```

### Project Structure

```bash
# View project structure
tree -I 'vendor|.git|bin'

# Run linter
golangci-lint run

# Format code
go fmt ./...
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Workflow

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes and add tests
4. Commit your changes: `git commit -m 'Add amazing feature'`
5. Push to the branch: `git push origin feature/amazing-feature`
6. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [age](https://github.com/FiloSottile/age) for modern encryption
- [BoltDB](https://github.com/etcd-io/bbolt) for embedded storage
- [Cobra](https://github.com/spf13/cobra) for CLI framework
- [promptui](https://github.com/manifoldco/promptui) for TUI components

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/imrancluster/sshmama/issues)
- **Discussions**: [GitHub Discussions](https://github.com/imrancluster/sshmama/discussions)
- **Wiki**: [GitHub Wiki](https://github.com/imrancluster/sshmama/wiki)

---

**Made with ❤️ by the sshmama community**

*Simplify your SSH workflow with secure, organized connection management.*
