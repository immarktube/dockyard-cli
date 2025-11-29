# dockyard exec Command

The `exec` command allows you to run **any arbitrary Git command** across all repositories defined in your Dockyard configuration.\
This is useful when you need full control and want to manually execute Git operations that are not covered by built-in Dockyard commands.

***

### 📌 **Usage**

```bash
dockyard exec [git arguments...]
```

* Everything after `exec` will be passed directly to `git`.
* Dockyard will automatically apply the command to **all repositories** in parallel.

***

### 📥 **Examples**

#### **1. Commit an empty commit across all repositories**

```bash
dockyard exec commit --allow-empty -m "Empty commit"
```

> ⚠️ On Windows CMD, you must escape flags like this:

```cmd
dockyard exec commit --% --allow-empty -m "Empty commit"
```

***

#### **2. Check the log in each repository**

```bash
dockyard exec log --oneline -5
```

***

#### **3. Create lightweight tags**

```bash
dockyard exec tag v1.0.0
```

***

#### **4. Clean local branches**

```bash
dockyard exec branch -D old-feature
```

***

#### **5. Stash local changes**

```bash
dockyard exec stash push -m "Temporary stash"
```

***

### 6. dummy push

```bash
dockyard exec commit --allowEmpty=true -m "dummy push"
```

***

### ⚠️ Behavior Notes

* Dockyard prints output per repository and does **not** stop when one repo fails.
* Works with your configured concurrency settings.
* Automatically respects any authentication hooks (e.g., token-injected remotes).
* Does _not_ run pre-defined Dockyard workflows — it simply runs the raw Git command.

***

### 🧠 Best Practices

* Use `exec` only when the built-in commands (`checkout`, `sync`, `push`, etc.) don’t provide what you need.
* Be careful with destructive operations like:
  * `reset --hard`
  * `clean -fd`
  * branch deletion
* For interactive commands (e.g., `git rebase -i`), usage may be limited or impractical.
