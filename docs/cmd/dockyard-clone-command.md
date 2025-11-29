# dockyard clone Command

The `dockyard clone` command allows you to clone **one or multiple Git repositories** into the directories defined in your Dockyard configuration file.

This command is useful when initializing a new development environment or onboarding new projects into your multi-repo workflow.

***

### 🚀 Usage

```bash
dockyard clone
```

Dockyard will read the repositories listed in your configuration file and clone each of them into the specified local path.

***

### 🛠️ Options

The `clone` command currently does not require additional flags.\
All repository URLs and target directories must be predefined in your Dockyard config file (e.g. `~/.dockyard/config.yaml`).

***

### 📁 Configuration Example

Here is an example of how repositories should be defined in the Dockyard config file:

```yaml
repositories:
  - name: service-account
    path: ~/workspace/service-account

  - name: service-order
    path: ~/workspace/service-order
```

***

### 💡 What Happens When You Run `dockyard clone`?

For each repository entry in your configuration:

1. Dockyard checks whether the target directory already exists
2.  If the directory **does not exist**, Dockyard performs:

    ```bash
    git clone <repo-url> <repo-path>
    ```
3. If the directory **already exists**, Dockyard prints a message and skips cloning\
   (to avoid overwriting your local work)

***

### 📝 Example Output

```
==> Cloning git repository service-account...
Cloned into /Users/mark/workspace/service-account

==> Cloning git repository service-order...
Directory already exists. Skipped.
```

***

### 🧩 Notes

* SSH or HTTPS cloning depends on the URL you specify in the config.
* If cloning fails (e.g., authentication issues), Dockyard will display the Git error.
* The command runs _sequentially_ unless you implement concurrency.
