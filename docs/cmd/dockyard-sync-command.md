# dockyard sync Command

### **Overview**

The `sync` command performs:

* `git fetch`
* `git pull`
* Optional rebase, autostash, or fast-forward actions

across **all repositories** defined in your Dockyard configuration.

This command ensures every repository is synchronized with its remote state, helping maintain consistent branches before applying changes, creating PRs, or performing batch operations.

***

### **Usage**

```bash
dockyard sync
```

Optionally:

```bash
dockyard sync --baseRef <branch>
```

***

### **Flags**

| Flag           | Description                                                                          |
| -------------- | ------------------------------------------------------------------------------------ |
| `--rebase`     | Use `git pull --rebase` instead of merge.                                            |
| `--autostash`  | Automatically stash local changes before rebasing (`git pull --rebase --autostash`). |
| `--ff-only`    | Only perform fast-forward merges; abort if a merge commit is required.               |
| `--dry-run`    | Show what would be executed without making changes.                                  |
| `-h`, `--help` | Show help information.                                                               |

_Note: Available flags may vary depending on your implementation._

***

### **Default Behavior**

* Runs `git fetch` for each repository.
* Runs `git pull` for each repository using:
  * **current branch's upstream** if no `--baseRef` is provided.
* Displays output grouped by repository.
* Executes operations concurrently using the configured concurrency limit.

***

### **Examples**

#### **1. Sync all repositories with their current upstream branches**

```bash
dockyard sync
```

Sample output:

```
==> Syncing services/auth
Fetching origin
Already up to date.

==> Syncing services/api
Fetching origin
Updating 91d2f..b7e4a
Fast-forward
```

***

#### **2. Sync all repositories to a specific base branch**

```bash
dockyard sync --baseRef main
```

Dockyard will:

* Checkout `main` (if not already on it)
* Fetch remote
* Pull from `origin/main`

***

#### **3. Sync using rebase strategy**

```bash
dockyard sync --rebase
```

Equivalent to running:

```
git pull --rebase
```

***

#### **4. Sync with auto-stash enabled**

```bash
dockyard sync --rebase --autostash
```

This is extremely convenient when you have local modifications but still want to update.

***

#### **5. Dry run without performing any fetch/pull**

```bash
dockyard sync --dry-run
```

Example output:

```
[DRY RUN] Would run: git fetch
[DRY RUN] Would run: git pull --rebase --autostash
```

***

### **Typical Use Cases**

* Before running batch commits (using `patch` or `updateYaml`)
* Before pushing changes (`dockyard push`)
* Before creating PRs (`dockyard createPR`)
* Ensuring all repos are aligned with `main` or `develop`
* Syncing microservice repos before global refactors or sweeping changes

***

### **Advanced Examples**

#### **6. Force all repos to sync with a specific branch regardless of their current branch**

```bash
dockyard sync --baseRef release/1.2
```

Dockyard will:

* Switch to `release/1.2`
* Fetch
* Pull

***

#### **7. Sync only using fast-forward merges**

```bash
dockyard sync --ff-only
```

If a repo requires a merge commit, sync will fail for that repo.

***

### **Common Errors & Troubleshooting**

#### **1. “No tracking information for the current branch”**

Occurs when the repo branch has **no upstream configured**.

Fix options:

*   Run:

    ```bash
    git branch --set-upstream-to=origin/<branch>
    ```
* Or ensure you always use `--baseRef main` when syncing new branches.

***

#### **2. Merge conflicts during sync**

When using merge:

```
Automatic merge failed; fix conflicts and commit the result.
```

Fix:

* Manually resolve the conflict in that repo.
* Re-run `dockyard sync`.

***

#### **3. Rebase conflicts**

When using `--rebase`:

```
CONFLICT (content): Merge conflict in file.yaml
```

Fix:

* Solve conflict
*   Continue rebase:

    ```
    git rebase --continue
    ```
* Then re-run `dockyard sync`.

***

#### **4. Local changes prevent pulling**

```
error: Your local changes to the following files would be overwritten...
```

Solutions:

* Use `--autostash`
* Or manually handle changes (`stash`, commit, reset, etc.)

***

### **Related Commands**

| Command                | When to use                                             |
| ---------------------- | ------------------------------------------------------- |
| `status`               | Check for uncommitted changes before syncing            |
| `exec`                 | Run custom git commands in bulk if `sync` is not enough |
| `push`                 | Push new changes after syncing                          |
| `checkout`             | Prepare branches before syncing                         |
| `patch` / `updateYaml` | Often run **after** `sync` to ensure updated base       |
