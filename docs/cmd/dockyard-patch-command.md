# dockyard patch Command

The `patch` command applies text-based modifications to a specific file across all repositories.\
It searches for a given regular expression and replaces the matched content with a user-defined string.

This is useful for batch-updating configuration files, version bumps, toggles, environment flags, YAML values, or any repetitive change that can be expressed with regex.

***

### 📌 **Usage**

```bash
dockyard patch --file <path> --old <regex> --new <string> [flags]
```

***

### 🧩 **Flags**

| Flag                | Description                                                           |
| ------------------- | --------------------------------------------------------------------- |
| `--file`            | Relative file path to patch (required).                               |
| `--old`             | Regex pattern to search for (required).                               |
| `--new`             | The replacement text (required). Supports `\n`, captured groups, etc. |
| `--dry-run`         | Performs a preview without modifying files.                           |
| `--message`         | Custom Git commit message. Optional.                                  |
| `--max-concurrency` | Limit parallel execution across repositories (optional).              |
| `--regex`           | Treat --old as regular expression                                     |

***

### ✅ **Basic Examples**

#### **1. Replace a version string**

```bash
dockyard patch \
  --file package.json \
  --old '"version": "1.0.0"' \
  --new '"version": "1.1.0"'
```

***

#### **2. Replace an environment toggle**

Switch all repos from staging → production:

```bash
dockyard patch \
  --file env.config \
  --old 'ENV=staging' \
  --new 'ENV=production'
```

***

#### **3. Add a new line after a match**

Adding structured YAML content:

```bash
dockyard patch \
  --file values.yaml \
  --old 'distroless: true' \
  --new 'distroless: true\n  GCP: true'
```

***

#### **4. Remove deprecated field**

```bash
dockyard patch \
  --file config.yaml \
  --old 'deprecatedField:.*' \
  --new ''
```

***

### 🔍 **Advanced Regex Examples**

#### **5. Replace only inside a specific block**

Match any `image:` inside a Deployment:

```bash
dockyard patch \
  --file deployment.yaml \
  --old 'image:.*' \
  --new 'image: myrepo/app:latest' \
  --regex
```

***

#### **6. Capture groups**

Convert:

```
timeout = 30s
```

to:

```
timeout = 45s
```

but only replace the number:

```bash
dockyard patch \
  --file config.ini \
  --old 'timeout = (\d+)s' \
  --new 'timeout = 45s' \
  --regex
```

***

#### **7. Multiline matching**

Replace an entire YAML section:

```bash
dockyard patch \
  --file config.yaml \
  --old 'resources:\n(\s*limits:.*\n\s*requests:.*)' \
  --new 'resources: {}' \
  --regex
```

***

### 🔧 **Dry Run Example**

Preview changes without modifying files:

```bash
dockyard patch \
  --file app.yaml \
  --old 'replicas: [0-9]+' \
  --new 'replicas: 3' \
  --regex \
  --dry-run
```

Dry run output example:

```
[DRY RUN] Would modify app.yaml in repo service-api
[DRY RUN] Would modify app.yaml in repo user-core
```

***

### 📝 **Commit Message Override**

```bash
dockyard patch \
  --file deployment.yaml \
  --old 'replicas: 1' \
  --new 'replicas: 3' \
  --message "Increase replicas to 3"
```

Otherwise Dockyard automatically commits with:

```
dockyard: patch file <filename>
```

***

### 🌍 **Real-world Examples**

#### **8. Switch container base image**

```bash
dockyard patch \
  --file Dockerfile \
  --old 'FROM node:16' \
  --new 'FROM node:20'
```

***

#### **9. Disable a feature flag**

```bash
dockyard patch \
  --file config.yaml \
  --old 'featureX: true' \
  --new 'featureX: false'
```

***

#### **10. Insert block into nginx.conf**

```bash
dockyard patch \
  --file nginx.conf \
  --old 'http {' \
  --new 'http {\n    include security.conf;'
```

***

#### **11. Update helm chart version**

```bash
dockyard patch \
  --file Chart.yaml \
  --old 'version: .*' \
  --new 'version: 2.0.0'
```

***

### ⚠️ Notes & Best Practices

* **Regex patterns must be quoted**—especially on Windows or zsh.
* Replacement text supports escaped sequences (`\n`, `\t`) because Dockyard interprets them in Go.
* For YAML-specific semantic updates, consider using `updateYaml` instead of `patch`.
* If the file does not exist, Dockyard will skip the repo automatically.
