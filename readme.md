# {Insert Name Here} - Your Local Development Environment Orchestrator

## Overview

{Insert Name Here} is a powerful CLI tool designed to enhance your local development workflow.T his tool automates and simplifies the process of managing and orchestrating multiple projects, making your development process more efficient.

## Features (To-Do List)
- [X] **Auto-Detect Projects**: Automatically detect projects in your local environment.
- [/] **Project Cloning**: Integrate with GitHub APIs to clone projects automatically.
- [X] **Proxy Management**: Check if the local proxy is running and manage its status.
- [ ] **CLI Dashboard**: A user-friendly command-line interface for managing local projects.
- [ ] **Easy Project Management**: Simplify the process of taking down and running projects.
- [ ] **Deployment Simplification**: Deploy projects easily with integrated commands.
- [ ] **GitHub App Integration**: Use a GitHub app for better management of your GitHub account.
- [ ] **Environment Orchestration**: Orchestrate your development environment with deployment capabilities.

## Getting Started
1. **Installation**: Ensure Go is installed on your system. Clone this repository and run `go build`.
2. **Basic Commands**(Coming Soon):

## Project Detection and serving
- The `start.go` file checks your proxy status and manages it accordingly.
- It scans the specified directory for projects containing a Makefile or Docker Compose file.
- You can select a project from the list, with priority given to `Makefile`, otherwise using `docker compose up -d`.

## Contribution
Your contributions are valuable! Feel free to fork, submit pull requests, or open issues for enhancements or bug fixes.

---

_Note: This README is subject to updates as the project progresses._


Dev notes/commands
build
```VERSION=$(git describe --tags)```
```go build -ldflags="-X 'main.version=$VERSION'" -o your_binary_name```
