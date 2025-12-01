# dockyard status Command

### **Overview**

The `status` command runs:

```
git status
```

inside **every managed repository** and prints the results in a clean, readable format.

This provides a quick overview of:

* Which repositories have unstaged changes
* Which repositories have staged but uncommitted changes
* Which repositories are clean
* Which repositories are on diverged branches
* Which repositories contain untracked files

It is one of the most commonly used commands for multi-repo workflows.

***

### **Usage**

```bash
dockyard status
```

No arguments are required or accepted.

***

### **Flags**

| Flag           | Description                                         |
| -------------- | --------------------------------------------------- |
| `--short`      | Show the short status format (`git status --short`) |
| `--branch`     | Show branch information (`git status --branch`)     |
| `--porcelain`  | Output in porcelain format for scripting            |
| `-h`, `--help` | Show help information                               |

_Note: Flags may vary depending on your implementation._

***

### **Default Behavior**

* Executes `git status` in each repository.
* Output is grouped per repository with clear headers.
* Stops at first error (e.g., invalid repo path or missing .git directory).

***

### **Examples**

#### **1. Check status for all repositories**

```bash
dockyard status
```

Sample output:

```
==> Repository: services/user-service
On branch main
Changes not staged for commit:
  modified: src/user.go

==> Repository: services/order-service
nothing to commit, working tree clean
```

***

#### **2. Show concise status (short format)**

```bash
dockyard status --short
```

Sample:

```
==> services/api
 M app/controllers/home.rb
?? tmp/log.txt
```

***

#### **3. Show branch & ahead/behind info**

```bash
dockyard status --branch
```

Sample:

```
## main...origin/main [ahead 1, behind 2]
 M src/main.go
```

***

#### **4. Machine-readable porcelain output**

Useful for automation scripts:

```bash
dockyard status --porcelain
```

***

### **Typical Use Cases**

* Quickly identifying modified repositories before creating a pull request.
* Checking which modules require commits or pushes.
* Validating repository cleanliness before running CI workflows.
* Inventorying untracked files or temporary changes.
* Daily routine check of all repo states in a monorepo-like environment.

***

### **Advanced Examples**

#### **5. Combine with external tooling (shell script)**

Monitor modified repos:

```bash
dockyard status --porcelain | grep -B1 " M "
```

***

#### **6. Run before a batch commit or patch**

```bash
dockyard status --short
dockyard patch ...
```

***

### **Common Errors & Troubleshooting**

#### **1. Repository not initialized**

```
fatal: not a git repository (or any of the parent directories): .git
```

Cause: The configured directory is not a valid Git repository.

Fix: Remove or correct the repo entry in your config.

***

#### **2. Permission issues**

```
error: cannot open .git/index
```

Fix: Check OS file permissions.

***

#### **3. Status appears different than expected**

Depending on your Git config, ignored files or submodules may behave differently.

Check:

```bash
git config --global status.showUntrackedFiles
git config --global core.ignorecase
```

***

### **Suggested Related Commands**

| Command    | Relationship                                                                |
| ---------- | --------------------------------------------------------------------------- |
| `sync`     | Useful to run before checking status to fetch/pull changes.                 |
| `exec`     | Can be used to run arbitrary git commands if status is insufficient.        |
| `patch`    | Often used after `status` to confirm and apply file modifications.          |
| `createPR` | Clean working trees (verified by `status`) are required before PR creation. |
