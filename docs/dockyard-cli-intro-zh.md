# 📦 Dockyard CLI

**Dockyard CLI** 是一个基于 Go 语言开发的命令行工具，旨在简化和自动化项目的构建、部署和管理流程。
它可以帮你管理成百上千个仓库的配置，这可以为您节省大量的时间和精力。

🔗 项目主页：[immarktube.github.io/dockyard-cli](https://immarktube.github.io/dockyard-cli/)

---

## 🚀 功能特性

- **模块化命令结构**：通过 `cmd/` 目录组织命令，便于扩展和维护。
- **配置驱动**：支持 `.dockyard.yaml` 配置文件，自定义构建和部署流程。
- **自动化执行**：内置任务执行器，自动处理常见的构建和部署任务。
- **易于集成**：可与现有的 CI/CD 流程无缝集成，提高开发效率。

---

## 🛠️ 安装与使用

### 安装

在以下页面下载最新版本的可执行文件并放置在与众多本地仓库同级目录： 
https://github.com/immarktube/dockyard-cli/releases
```
/your-workspace/
├── dockyard-cli         # Dockyard CLI 可执行文件（需放在此处）
├── kubernetesDemo       # 仓库1
├── careeranalyse-web    # 仓库2
├── readList             # 仓库3
```

### 使用

1. 在项目根目录创建 `.dockyard.yaml` 配置文件，定义构建和部署任务。
2. 运行以下命令执行任务：

```bash
dockyard --help
```

详细的使用指南请参考：[Dockyard CLI 使用指南](https://github.com/immarktube/dockyard-cli/wiki)

---

## 📁 项目结构

```
dockyard-cli/
├── cmd/             # 命令定义
├── command/         # 命令实现
├── config/          # 配置解析
├── docs/            # 文档
├── executor/        # 任务执行器
├── utils/           # 工具函数
├── .dockyard.yaml   # 示例配置文件
├── .env             # 示例配置文件
├── main.go          # 主程序入口
└── build.sh         # 构建脚本
```

---

## 📄 示例配置 `.dockyard.yaml`

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
提示: 你一样可以将token直接定义在 **.dockyard.yaml** 中，但这会直接暴露你的敏感信息在配置文件中。

---

## 🤝 贡献指南

欢迎贡献代码、提交问题或提出改进建议：

1. Fork 本仓库。
2. 创建新分支进行开发。
3. 提交 Pull Request。

---

## 📄 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](https://github.com/immarktube/dockyard-cli/blob/main/LICENSE) 文件。
