# dockyard complete Command

The `dockyard complete` command is used to **generate shell completion scripts** for Dockyard.\
These scripts enable features like:

* **Tab autocompletion** for commands
* **Flag suggestions**
* **Subcommand hints**

Dockyard supports generating completion scripts for commonly used shells such as:

* Bash
* Zsh
* Fish
* PowerShell

***

### 🚀 Usage

```bash
dockyard complete [shell]
```

`[shell]` indicates the type of shell you want completion output for.

***

### 🛠️ Arguments

| Argument | Required | Description                                           |
| -------- | -------- | ----------------------------------------------------- |
| `shell`  | Yes      | The shell type (`bash`, `zsh`, `fish`, `powershell`). |

***

### 🎯 Example Usage

#### **Bash**

```bash
dockyard complete bash > /etc/bash_completion.d/dockyard
```

or for local user installation:

```bash
dockyard complete bash > ~/.local/share/bash-completion/dockyard
```

Reload:

```bash
source ~/.bashrc
```

***

#### **Zsh**

```bash
dockyard complete zsh > "${fpath[1]}/_dockyard"
```

Then reload:

```bash
source ~/.zshrc
```

***

#### **Fish**

```bash
dockyard complete fish > ~/.config/fish/completions/dockyard.fish
```

***

#### **PowerShell**

```powershell
dockyard complete powershell | Out-String | Invoke-Expression
```

To persist:

```powershell
dockyard complete powershell > dockyard.ps1
```

***

### 📌 Notes

* This command **prints** the completion script to stdout.
* You must redirect (`>`) the output to the correct file.
* After installation, restart your shell or reload your config file.
