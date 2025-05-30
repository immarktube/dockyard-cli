# 📦 Dockyard CLI

**Dockyard CLI** is a command-line tool built in Go to simplify and automate project build, deployment, and task execution.

🔗 Project Homepage: [immarktube.github.io/dockyard-cli](https://immarktube.github.io/dockyard-cli/)

---

## 🚀 Features

- **Modular Command Structure**: Organized via the `cmd/` directory for easy extension and maintenance.
- **Configuration-Driven**: Supports `.dockyard.yaml` for defining custom build and deployment pipelines.
- **Automated Task Execution**: Built-in task runner for handling common project workflows.
- **CI/CD Friendly**: Easily integrates into your existing automation pipelines.

---

## 🛠️ Installation & Usage

### Installation

Ensure you have Go (version 1.16 or later) installed:

```bash
go install github.com/immarktube/dockyard-cli@latest
```

### Usage

1. Create a `.dockyard.yaml` file at your project root to define tasks.
2. Run your tasks using:

```bash
dockyard-cli run
```

For detailed usage instructions, visit: [Dockyard CLI Documentation](https://immarktube.github.io/dockyard-cli/)

---

## 📁 Project Structure

```
dockyard-cli/
├── cmd/             # Command definitions
├── command/         # Command implementations
├── config/          # Configuration parsing
├── docs/            # Documentation
├── executor/        # Task runner
├── utils/           # Utility functions
├── .dockyard.yaml   # Example config file
├── main.go          # Entry point
└── build.sh         # Build script
```

---

## 📄 Example `.dockyard.yaml`

```yaml
build:
  steps:
    - name: Build the project
      command: go build -o bin/app main.go
    - name: Run tests
      command: go test ./...
deploy:
  steps:
    - name: Deploy to server
      command: scp bin/app user@server:/path/to/deploy
```

---

## 🤝 Contributing

We welcome contributions, issue reports, and suggestions!

1. Fork this repository.
2. Create a new feature branch.
3. Submit a Pull Request.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/immarktube/dockyard-cli/blob/main/LICENSE) file for details.
