# `updateYaml` Command Documentation

## Overview

The `updateYaml` command modifies a specific YAML file across **all configured repositories** by updating a given node path with a new value.

It is designed for **batch configuration updates**, especially useful when maintaining consistent YAML configurations (e.g., Kubernetes manifests, Helm values, environment configs) across multiple repositories.

---

## Usage

```bash
dockyard updateYaml [flags]
````

---

## Flags

| Flag                | Description                                                                   |
| ------------------- | ----------------------------------------------------------------------------- |
| `--filePath <path>` | Relative path to the YAML file (must end with `.yaml` or `.yml`)              |
| `--nodePath <path>` | Dot-separated path to the target node (e.g., `spec.template.spec.containers`) |
| `--value <value>`   | Value to set at the specified node path                                       |
| `--createIfAbsent`  | Automatically create missing nodes if they do not exist                       |
| `--dry-run`         | Preview changes without modifying files                                       |
| `-h, --help`        | Show help information                                                         |

---

## Example

```bash
dockyard updateYaml \
  --filePath config/app.yaml \
  --nodePath metadata.name \
  --value my-app \
  --createIfAbsent \
  --message "Update app name"
```

---

## Behavior

For each repository, the command performs:

1. Resolve the target file path:

   ```
   <repo.Path>/<filePath>
   ```

2. Parse the YAML file

3. Update the specified node path:

    * If `--createIfAbsent` is enabled, missing nodes will be created
    * Otherwise, the operation fails if the path does not exist

4. Write changes back to the file

5. Print result for each repository

---

## Node Path Format

The `--nodePath` uses **dot notation** to locate the YAML node.

### Example YAML:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
```

### Example command:

```bash
dockyard updateYaml \
  --filePath deployment.yaml \
  --nodePath metadata.name \
  --value new-name
```

👉 Result:

```yaml
metadata:
  name: new-name
```

---

## Behavior with Missing Nodes

### ❌ Without `--createIfAbsent`

If the node path does not exist:

```text
failed to update path 'metadata.labels.app': path 'metadata.labels' is not a mapping node
```

---

### ✅ With `--createIfAbsent`

```bash
dockyard updateYaml \
  --filePath deployment.yaml \
  --nodePath metadata.labels.app \
  --value my-app \
  --createIfAbsent
```

👉 Result:

```yaml
metadata:
  labels:
    app: my-app
```

---

## Dry Run

```bash
dockyard updateYaml \
  --filePath config.yaml \
  --nodePath spec.replicas \
  --value 3 \
  --dry-run
```

Output:

```text
📝 Dry-run: Would modify /repo/path/config.yaml
```

---

## Multi-Repository Behavior

* The command runs concurrently across all repositories
* Each repository is processed independently
* If a file does not exist in a repo, it will fail for that repo only

Example:

```text
repoA: ✅ UpdateYaml complete!
repoB: ❌ Failed to write config.yaml: file not found
```

---

## Validation Rules

* `--filePath` must end with `.yaml` or `.yml`
* `--nodePath` must be a valid dot-separated path
* `--value` is treated as a string (unless your implementation parses types)

---

## Common Errors & Solutions

---

### ❌ File path invalid

```text
file path must end with .yml or .yaml
```

**Fix:**

```bash
--filePath config.yaml
```

---

### ❌ Path not found

```text
path '' is not a mapping node
```

**Cause:**

* Intermediate node is not a map
* Path structure is incorrect

**Fix:**

* Verify YAML structure
* Use `--createIfAbsent`

---

### ❌ YAML structure issues

If YAML contains unexpected formats (e.g., arrays instead of maps), path resolution may fail.

---

## Best Practices

---

### ✅ Use `--dry-run` before applying changes

```bash
dockyard updateYaml --dry-run ...
```

---

### ✅ Use `--createIfAbsent` for safe automation

Prevents failures when nodes are missing.

---

### ✅ Keep node paths simple and explicit

Avoid overly deep or ambiguous paths.

---

## Limitations

* Only supports **map-based YAML paths** (no array indexing like `[0]`)
* Value is applied as-is (no advanced type inference unless implemented)
* Does not automatically commit changes (you can use `dockyard apply` afterward)

---

## Suggested Workflow

```bash
dockyard updateYaml --filePath config.yaml --nodePath spec.replicas --value 3
dockyard apply --include "config.yaml" -m "Update replicas"
```

---

## Summary

The `updateYaml` command enables:

* ✅ Batch YAML updates across repositories
* ✅ Structured node-level modification
* ✅ Optional auto-creation of missing nodes
* ✅ Safe preview via dry-run
* ✅ Concurrent execution

---

## Philosophy

> **Automate repetitive config changes safely and consistently**

This command helps enforce configuration consistency across multiple repositories with minimal manual effort.

---

```
```
