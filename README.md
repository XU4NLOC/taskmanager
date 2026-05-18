# taskmanager

A fast, lightweight CLI task manager built with Go and Cobra. Manage your to-do list directly from the terminal — no browser, no app, just your keyboard.

---

## Features

- Add tasks with a single command
- List all tasks with their status
- Mark tasks as complete
- Delete tasks by ID
- Tasks persist across sessions via a local JSON file (`~/.tasks.json`)

---

## Installation

**Requirements:** Go 1.21+

```bash
git clone https://github.com/yourusername/taskmanager.git
cd taskmanager
go install .
```

Make sure your Go binary path is in your shell's PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## Usage

```bash
# Add a new task
taskmanager add "Buy groceries"

# List all tasks
taskmanager list

# Mark a task as complete
taskmanager complete 1

# Delete a task
taskmanager delete 1
```

---

## Project Structure

```
taskmanager/
├── main.go           # Entry point
├── cmd/
│   ├── root.go       # Root Cobra command
│   ├── add.go        # add command
│   ├── list.go       # list command
│   ├── complete.go   # complete command
│   └── delete.go     # delete command
└── storage/
    └── storage.go    # JSON persistence layer
```

---

## Data Storage

Tasks are stored locally at `~/.tasks.json` in a human-readable format:

```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "done": false,
    "created_at": "2026-05-18 10:30:00"
  }
]
```

---

## Built With

- [Go](https://golang.org/) — systems language
- [Cobra](https://github.com/spf13/cobra) — CLI framework (used by Docker, Kubernetes, Hugo)

---

## Motivation

This project is part of my journey to master Go and transition into backend/systems engineering. It covers core Go concepts including structs, interfaces, file I/O, JSON encoding, and CLI design patterns.
# taskmanager
