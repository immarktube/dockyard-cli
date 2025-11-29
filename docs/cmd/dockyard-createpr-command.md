# dockyard createPR Command

The `dockyard createPR` command creates Pull Requests across all repositories managed by Dockyard.\
It is designed for batch PR creation after automated updates—such as YAML modifications, config updates, template propagation, or branch syncing.

***

### 🚀 Usage

```bash
dockyard createPR --title "<title>" --body "<body>"
```

***

### 🛠️ Flags

| Flag      | Type   | Required | Description                               |
| --------- | ------ | -------- | ----------------------------------------- |
| `--title` | string | Yes      | The title of the pull request.            |
| `--body`  | string | No       | The description/body of the pull request. |

***

### 📌 Behavior

* Automatically opens a Pull Request for **each repository** listed in `dockyard.yaml`.
* Uses the repository’s authenticated GitHub token to perform GitHub API calls.
* If the head branch does not exist on GitHub, PR creation will fail.
* If the PR already exists, the command handles or reports the duplication.

***

### 🎯 Example Usage

#### Create a PR for all repositories

```bash
dockyard createPR --title "Update configs" --body "Automated config changes"
```

Creates a PR similar to:

```
✅ PR created successfully for repo-A [branch-name]: htts://www.github.com/1234
✅ PR created successfully for repo-B [branch-name]: htts://www.github.com/1234
✅ PR created successfully for repo-C [branch-name]: htts://www.github.com/1234
```

***

### 📝 Notes

* Ensure that:
  * The head branch is **pushed** to GitHub before creating PRs.
  * Authentication tokens in your Dockyard config have **repo** permissions.
* Works seamlessly after:
  * `dockyard updateYaml`
  * `dockyard copyFile`
  * `dockyard exec` batch commits
  * `dockyard sync`
