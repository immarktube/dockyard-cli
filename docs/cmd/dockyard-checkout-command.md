# checkout Command Guide

### 📌 Overview

`dockyard checkout` is used to batch switch or create branches across multiple repositories defined in your `.dockyard.yaml` configuration.

This command allows you to:

* Switch to an existing branch in each repository
* Create a new branch based on:
  * A remote branch (default: `master`)
  * A specified base branch
  * A tag
  * A commit hash
* Respect per-repository base reference rules defined in config
* Provide a consistent new branch name across all repositories

***

### 🧠 How It Works

When running:

```sh
dockyard checkout <branch-name>
```

For each repository, Dockyard will:

1. Attempt to switch to the local branch
2. If missing, try to create a new branch based on:
   * `baseRef` defined in `.dockyard.yaml` (optional)
   * Otherwise default to `master`
3. `baseRef` can be:
   * A remote branch (e.g., `develop`, `prod`)
   * A tag (e.g., `v1.2.0`)
   * A commit hash (e.g., `a13b91c`)
4. Create the branch without establishing a remote tracking branch
5. Leave the branch ready for future `sync`, `push`, and PR creation

***

### 🛠 Example Usages

#### ✔ Create or switch to a branch named `feature/payment` on all repos

```sh
dockyard checkout feature/payment
```

#### ✔ Repositories with custom base references

In `.dockyard.yaml`:

```yaml
repositories:
  - path: service-user
    baseRef: develop

  - path: service-billing
    baseRef: v1.4.0  # based on tag

  - path: service-order
    baseRef: 4d23af1 # based on commit hash

  - path: service-gateway
    # no baseRef → defaults to "master"
```

Running:

```sh
dockyard checkout feature/payment
```

Result:

| Repository      | Base Reference   | Action                                     |
| --------------- | ---------------- | ------------------------------------------ |
| service-user    | develop          | Create branch from remote branch if needed |
| service-billing | tag v1.4.0       | Create branch based on tag                 |
| service-order   | commit `4d23af1` | Create branch from commit                  |
| service-gateway | master           | Default behavior                           |

***

### ⚠️ Notes & Behavior Details

#### Branch Tracking

Branches created from tags or commit hashes **do not track any remote branch**.\
This is expected and safe — future operations such as:

* `dockyard sync`
* `dockyard push`
* PR creation

will continue to work normally as long as the branch exists locally.

#### Push Behavior

If created from tag or commit hash:

* `git push origin HEAD` works normally
* The remote branch will be created automatically

#### Error handling

If a baseRef is invalid (e.g., nonexistent tag or commit), Dockyard will print an error for that repo but continue processing other repos.

***

### 💡 Recommendation

Use per-repository `baseRef` to keep full flexibility:

* Most repos → base from `master`
* Exception repos → base from tag or commit hash

Example:

```yaml
baseRef: master
```
