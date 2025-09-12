# 📦 Dockyard CLI 

**Dockyard CLI** is a command-line tool built in Go to simplify and automate project build, deployment, and task execution.
It helps you manage configurations across hundreds of repositories, saving you significant time and effort.

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
Download the latest release from below link and place the executable alongside your local repositories:  
https://github.com/immarktube/dockyard-cli/releases
```text
/your-workspace/
├── dockyard-cli         # Dockyard CLI 可执行文件（需放在此处）
├── kubernetesDemo       # 仓库1
├── careeranalyse-web    # 仓库2
├── readList             # 仓库3
```

### Usage

1. Create a `.dockyard.yaml` file at your project root to define tasks.
2. Run your tasks using:

```bash
dockyard --help
```

For detailed usage instructions, visit: [Dockyard CLI Documentation](https://github.com/immarktube/dockyard-cli/wiki)

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
├── .env             # Example config file
├── main.go          # Entry point
└── build.sh         # Build script
```

---

## 📄 Example `.dockyard.yaml`

```yaml
global:
  owner: immarktube
  authToken: ${GITHUB_TOKEN}
  apiBaseURL: https://api.github.com
  gitBaseURL: https://github.com
  concurrency: 5
  noHook: true

repositories:
  - path: kubernetesDemo
    baseRef: fb6512a5b8a5b763e0b2e8634bad4cd713239c48
  - path: careeranalyse-web
    baseRef: 1.0.0
  - path: readList
    baseRef: master

hook:
  pre: echo "Running pre hook"
  post: echo "Running post hook"
```

## 📄 Example `.env`

```env
GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```
Note: you can also define the token in **.dockyard.yaml**, but this will appear your sensitive data in config file directly.

---

## 🤝 Contributing

We welcome contributions, issue reports, and suggestions!

1. Fork this repository.
2. Create a new feature branch.
3. Submit a Pull Request.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/immarktube/dockyard-cli/blob/main/LICENSE) file for details.
