# dockyard push Command

### **Overview**

The `push` command runs `git push` across **all managed repositories**.\
It is typically used after performing batch updates using commands such as `patch`, `updateYaml`, `copyFile`, or `exec` (with commit operations).

This command ensures that all local commits are propagated to their corre谢·sponding remote branches.

***

### **Usage**

```bash
dockyard push 
```

***

### **Default Behavior**

If no flags are provided:

* The tool will push the **current branch** of each repository.
* The push is equivalent to running:

```bash
git push origin current-branch
```

for every repo.

***

### **Examples**

#### **1. Push the current branch in all repositories**

```bash
dockyard push
```

This pushes the current active branch in each repository to the remote `origin`.

***

#### **2. Force push all repositories**

```bash
dockyard push --force
```

Runs:

```bash
git push --force
```

⚠️ **Warning:** This overwrites remote history.\
Use only if you know what you're doing.

***

### **Typical Workflow Example**

```bash
dockyard patch --file config/app.yaml --old 123 --new 456 --message '123456'
dockyard push
```

This batch-updates a file and pushes all resulting commits.

***

### **Error Handling**

#### **1. Remote branch does not exist**

```
error: src refspec my-branch does not match any
```

Cause: The branch does not exist in that repo.\
Solution: Ensure the branch is created before pushing.

***

#### **2. Rejected non-fast-forward**

```
! [rejected]        main -> main (non-fast-forward)
```

Solutions:

*   Pull latest changes:

    <pre class="language-bash"><code class="lang-bash">dockyard sync
    <strong>dockyard push
    </strong></code></pre>
*   Or force push (dangerous):

    ```bash
    dockyard push --force
    ```

***

#### **3. Repository has no commits**

```
fatal: No commits yet
```

Solution: Commit first:

```bash
dockyard exec -- git commit --allow-empty -m "Initial commit"
dockyard push
```
