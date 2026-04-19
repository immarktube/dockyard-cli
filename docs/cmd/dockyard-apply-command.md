# `apply` Command Documentation

## Overview

The `apply` command stages selected files and creates a commit across **all configured repositories**.

It is designed to provide a **safe and flexible alternative** to manually running `git add` and `git commit` in multiple repositories.

Unlike a plain commit, `apply` requires you to explicitly define which files should be included, helping prevent accidental commits of unwanted changes.

---

## Usage

```bash
dockyard apply [flags]
````

---

## Flags

| Flag                   | Description                                         |
| ---------------------- | --------------------------------------------------- |
| `-m, --message <msg>`  | Commit message                                      |
| `--include <patterns>` | Files or glob patterns to include (comma-separated) |
| `--exclude <patterns>` | Files or glob patterns to exclude (comma-separated) |
| `--all`                | Include all files (`git add .`)                     |
| `--dry-run`            | Preview changes without applying them               |
| `-h, --help`           | Show help information                               |

---

## Important Rules

### ✅ You must specify one of:

* `--include`
* `--all`

Otherwise, the command will do nothing.

---

### ⚠️ `--exclude` cannot be used alone

This is **invalid**:

```bash
dockyard apply --exclude "*.env" -m "msg"
```

You must combine it with:

* `--include`, or
* `--all`

---

## Examples

---

### 1. Commit specific files

```bash
dockyard apply -m "update config" --include "config.yaml"
```

---

### 2. Commit multiple files

```bash
dockyard apply -m "update files" --include "a.yaml,b.yaml,c.json"
```

---

### 3. Commit using glob patterns

```bash
dockyard apply -m "update yaml files" --include "*.yaml"
```

---

### 4. Include + exclude

```bash
dockyard apply -m "update configs" \
  --include "*.yaml" \
  --exclude "secret.yaml"
```

---

### 5. Commit all files

```bash
dockyard apply -m "update all" --all
```

---

### 6. Commit all except some files

```bash
dockyard apply -m "update all except env" \
  --all \
  --exclude "*.env"
```

---

### 7. Dry run (preview only)

```bash
dockyard apply -m "update yaml" \
  --include "*.yaml" \
  --dry-run
```

Example output:

```
[DRY-RUN] repoA -> add: [config.yaml values.yaml]
[DRY-RUN] repoB -> add: [deployment.yaml]
```

---

## Behavior

For each repository, the command performs:

1. Resolve files using `--include` or `--all`
2. Run `git add` on matched files
3. Apply `--exclude` (if provided) using `git reset`
4. Run `git commit -m <message>`

---

## File Matching

* Supports **glob patterns** (e.g., `*.yaml`, `config/*.json`)
* Multiple patterns can be separated by commas
* Matching is **case-sensitive** (important!)

### Example:

```bash
--include "README.md"   ✅ matches
--include "readme.md"   ❌ does NOT match
```

---

## Multi-Repository Behavior

* Each repository is processed independently
* If a file does not exist in a repo, it is skipped
* If no files match in a repo, it will be skipped

Example:

```
repoA: ✅ committed config.yaml
repoB: ⚠️ config.yaml not found, skipped
```

---

## Common Errors & Solutions

---

### ❌ Nothing to commit

```
nothing to commit, working tree clean
```

**Cause:**

* No files matched `--include`
* Files were not staged

**Solution:**

* Check file names and patterns
* Use `--all` if needed

---

### ❌ File not matched (case issue)

```
--include "readme.md"
```

But actual file:

```
README.md
```

**Fix:**

```bash
dockyard apply --include "README.md"
```

---

### ❌ Unintended files not committed

**Cause:**

* Files were not included in `--include`

**Fix:**

* Expand include pattern:

```bash
--include "*.md,*.json"
```

---

## Best Practices

---

### ✅ Use include for precise control

```bash
dockyard apply -m "update yaml" --include "*.yaml"
```

---

### ✅ Use dry-run before large operations

```bash
dockyard apply --include "*.yaml" --dry-run
```

---

### ⚠️ Be careful with `--all`

```bash
dockyard apply --all -m "update"
```

This may include:

* unintended files
* logs
* temporary files

---

## Summary

The `apply` command provides:

* ✅ Safe batch commits across repositories
* ✅ Fine-grained file selection
* ✅ Support for glob patterns
* ✅ Optional exclusions
* ✅ Dry-run preview
* ✅ Concurrent execution

---

## Philosophy

> **Explicit is better than implicit**

`apply` only commits what you explicitly select, preventing accidental changes across multiple repositories.

---

```
```
