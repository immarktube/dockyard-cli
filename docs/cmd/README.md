# Dockyard CLI — Commands Index

## Dockyard CLI Commands Overview

A quick navigation guide to all available Dockyard CLI commands.

***

### 📁 Repository Operations

#### [**`checkout`**](dockyard-checkout-command.md)

Batch checkout branch in all repositories.

#### [**`clone`**](dockyard-clone-command.md)

Run `git clone` across all repositories.

#### [**`sync`**](dockyard-sync-command.md)

Run `git fetch` and `git pull` across all repositories.

#### [**`status`**](dockyard-status-command.md)

Run `git status` across all repositories.

#### [**`apply`**](dockyard-apply-command.md)
Apply all pending changes in the current branch across all repositories.

#### [**`push`**](dockyard-push-command.md)

Run `git push` across all repositories.

#### [**`exec`**](dockyard-exec-command.md)

Run arbitrary git command across all repositories.

#### [**`run`**](dockyard-run-command.md)

Run arbitrary shell command in all repositories.

***

### 📝 File & Config Operations

#### [**`copyFile`**](dockyard-copyfile-command.md)

Copy a file from one path to another inside each repository.

#### [**`patch`**](dockyard-patch-command.md)

Modify a specific file in all repositories and commit the change.

#### [**`updateYaml`**](dockyard-updateyaml-command.md)

Modify a specific YAML file in all repositories and commit the change.

***

### 🔀 Pull Request Operations

#### [**`createPR`**](dockyard-createpr-command.md)

Create pull requests for all modified repositories.

***

### 🧰 Utility

#### [**`completion`**](dockyard-completion-command.md)

Generate the autocompletion script for the specified shell.

#### [**`help`**](dockyard-help-command.md)

Help about any command.

