# dockyard copyFile Command

The `dockyard copyFile` command copies a file from the each repository folder into every repository another/new folder managed by Dockyard.\
It is useful when you need to distribute shared configuration files, templates, scripts, or other project assets across multiple repos.

***

### 🚀 Usage

```bash
dockyard copyFile --source <source-file> --target <relative-path> [--dry-run]
```

***

### 🛠️ Flags

<table><thead><tr><th>Flag</th><th>Type</th><th>Required</th><th>Description</th></tr></thead><tbody><tr><td><code>--source</code></td><td>string</td><td>Yes</td><td>Path to the repository relative path you want to copy.</td></tr><tr><td><code>--target</code></td><td>string</td><td>Yes</td><td>Destination <strong>relative path inside each repository</strong>. Existing files will be overwritten.</td></tr><tr><td><code>--dry-run</code></td><td>bool</td><td>No</td><td>If set, shows what would happen without modifying any repository.</td></tr><tr><td><p></p><pre><code>--message
</code></pre></td><td>string</td><td>yes</td><td>git commit message</td></tr></tbody></table>

***

### 📌 Behavior

* The source file is **copied** from each repository specific folder to each repository listed in `dockyard.yaml`.
* The source path and target path is interpreted **relative to each repo root**.
* Directories in the target path are created automatically if they do not exist.
* If the file already exists in a repo, it will be overwritten unless you implement additional options.

***

### 🎯 Example Usage

#### Copy a script into all repositories

```bash
dockyard copyFile --source ./scripts/setup.sh --target scripts/setup.sh --message commit setup.sh
```

This places `setup.sh` into:

```
repo-A/scripts/setup.sh
repo-B/scripts/setup.sh
repo-C/scripts/setup.sh
...
```

***

#### Copy a config file with dry-run mode

```bash
dockyard copyFile --source .editorconfig --target .editorconfig --dry-run
```

Output example:

```
[DRY-RUN] Would copy .editorconfig to repo-A/.editorconfig
[DRY-RUN] Would copy .editorconfig to repo-B/.editorconfig
...
```

No file operations are performed.

***

#### Copy a YAML template into each repo

```bash
dockyard copyFile --source template.yaml --target configs/template.yaml --message 'commit update'
```

If `configs/` does not exist inside a repo, it will be created.

***

### 📝 Notes

* `copyFile`  commit changes automatically.
* Works well together with sync, exec, and updateYaml workflows.
