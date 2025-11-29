# Dockyard CLI — Commands Index

Synopsis\
A central index page listing available commands, global options, directory conventions, and examples. Each command should have its own detail page under `docs/cmd/`.

Prerequisites

* Compiled executable placed as a sibling to your local repositories (see Directory Structure).
* Each repository listed in the tool configuration is available locally.
* `git` installed and repositories are in a committable state.

Global notes

* Executable names: `dockyard-cli` or `dockyard`
* Supported project types: Go, Java (Spring Boot / Maven), TypeScript / JavaScript (npm), Vue, etc.
* Configuration: tool reads repository list and concurrency settings from project configuration (see code for details).

Global flags / environment variables Common options can be supplied as flags or via environment variables:

* `--max-concurrency` the maximum number of concurrent operations (integer) defaults to 5
* `--no-hook` (boolean) disable pre-operation and post-operation hooks

Commands index

1. `checkout` - Checkout a specific branch or commit in each repository.
2. `clone` - Clone repositories listed in configuration.
3. `complete` - complete command placeholder.
4. `copyFile` - Copy a file to each repository and optionally commit.
5. `createPR` - Create pull requests for changes in each repository.
6. `exec` - Execute a git command across configured repositories.
7. `help` - Show help for Dockyard CLI or a specific command.
8. `patch` - Modify a specific file in all repositories and commit the change.
9. `push` - Push committed changes to remote repositories for all the configured repositories.
10. `run` - Run a shell command in each repository.
11. `status` - Show git status for each repository.
12. `sync` - Run 'git fetch' and 'git pull' across all repositories
13. `updateYaml` - Update a YAML file across configured repositories and optionally commit.

Documentation maintenance

* Keep one detailed page per command under `docs/cmd/`, e.g.:
  * `docs/cmd/updateYaml.md`
  * `docs/cmd/listRepos.md`
* Keep `docs/cmd/cmdRoot.md` as a concise index linking to those pages.
