# dockyard run Command

### **Overview**

The `run` command executes an **arbitrary shell command** inside **each repository** managed by the tool.

It differs from the `exec` command in that:

* **`exec` runs arbitrary&#x20;**_**git**_**&#x20;commands** (`git status`, `git fetch`, etc.)
* **`run` runs arbitrary&#x20;**_**shell**_**&#x20;commands** (`ls`, `npm install`, `go build`, `sed`, etc.)

This makes `run` useful for performing non-git batch operations such as building, testing, installing dependencies, or modifying files across multiple repositories.

***

### **Usage**

```bash
dockyard run -- <command> [args...]
```

> Important: A `--` is usually required to separate tool arguments from the command you want to run.

***

### **Flags**

| Flag                  | Description                                                                                      |
| --------------------- | ------------------------------------------------------------------------------------------------ |
| `--shell <path>`      | Use a custom shell to run the command (default depends on OS: `bash` on Unix, `cmd` on Windows). |
| `--continue-on-error` | Continue running the command in other repositories even if it fails in one.                      |
| `--dry-run`           | Show the command that _would_ be executed, but do not run it.                                    |

***

### **Default Behavior**

* The specified command is executed inside each repository's root directory.
* Stops on the first error (unless `--continue-on-error` is set).
* Output from each repository is printed with repo name prefixes for clarity.

***

### **Examples**

#### **1. Run `ls` in every repository**

```bash
dockyard run -- ls -la
```

Runs `ls -la` in the root of every managed repo.

***

#### **2. Install dependencies across all Node.js projects**

```bash
dockyard run -- npm install
```

***

#### **3. Run a Go build in every Go repository**

```bash
dockyard run -- go build ./...
```

***

#### **4. Run a custom script file**

```bash
dockyard run -- ./scripts/setup.sh
```

***

#### **5. Using a custom shell**

```bash
dockyard run --shell /bin/zsh -- echo "Hello from zsh"
```

***

#### **6. Run commands on Windows using PowerShell**

```bash
dockyard run --shell powershell -- Get-ChildItem
```

Or with default shell (cmd):

```bash
dockyard run -- dir
```

***

#### **7. Continue even if some repositories fail**

```bash
dockyard run --continue-on-error -- go test ./...
```

***

#### **8. Dry-run (show but do not execute)**

```bash
dockyard run --dry-run -- yarn build
```

***

#### **9. Use run to perform batch text replacement (cross-platform example)**

**Linux/macOS**

```bash
dockyard run -- sed -i '' 's/oldValue/newValue/g' config.yaml
```

**Windows cmd**

```bash
dockyard run -- powershell -Command "(Get-Content config.yaml).replace('old','new') | Set-Content config.yaml"
```

***

### **Advanced Examples**

#### **10. Execute multiple commands using a single shell call**

```bash
dockyard run -- bash -c "npm install && npm test"
```

***

#### **11. Clean build artifacts in all repos**

```bash
dockyard run -- rm -rf build dist out
```

***

#### **12. Find large files in each repository**

```bash
dockyard run -- find . -size +50M
```

***

#### **13. Batch format code (Go example)**

```bash
dockyard run -- go fmt ./...
```

***

#### **14. Run Python script in all repositories**

```bash
dockyard run -- python tools/check_config.py
```

***

### **Common Errors & Troubleshooting**

#### **1. Command not found**

```
exec: "npm": executable file not found
```

Solution: Ensure the command exists in PATH.

***

#### **2. Shell syntax errors on Windows**

If the command works on Linux but fails on Windows:

Try:

```bash
dockyard run --shell powershell -- <command>
```

or

```bash
dockyard run --shell bash -- <command>
```

(if using Git Bash)

***

#### **3. Tool stops at first failing repo**

If one repository fails, the command stops:

```
Command failed in repo A
Stopping execution
```

Solution:

```bash
dockyard run --continue-on-error -- <command>
```

***

### **Suggested Use Cases**

* Building or testing across many repos.
* Running linters or static analysis tools.
* Cleaning or preparing environments.
* Applying consistent file operations.
* Checking disk usage or debugging.
* Running scripts that automate bulk tasks.
