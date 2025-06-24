# go-todo-cli

A simple and efficient CLI-based Todo application written in Go. Manage your tasks directly from the terminal.

## Features
- Add todos with an optional due date (default: 24h from now)
- View all todos in a table, sorted by due date
- View a todo by its ID
- Update the status of a todo (pending, in-progress, completed)
- Delete a todo by its ID

## Installation
1. Clone the repository:
   ```sh
   git clone <repo-url>
   cd go-todo-cli
   ```
2. Build the app:
   ```sh
   go build -o go-todo
   ```
3. Move the binary to your PATH:
   ```sh
   sudo mv go-todo /usr/local/bin/
   ```

## Usage
- Add a todo:
  ```sh
  go-todo add "Task title" --due 2025-06-25
  ```
- View all todos:
  ```sh
  go-todo view
  ```
- View a todo by ID:
  ```sh
  go-todo get 1
  ```
- Update status:
  ```sh
  go-todo update 1 completed
  ```
- Delete a todo:
  ```sh
  go-todo delete 1
  ```

## Status Options
- pending
- in-progress
- completed

## Data Storage
Todos are stored in a local `todos.json` file.

