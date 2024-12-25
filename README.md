# Task Tracker

A simple task tracker command-line application written in Go. It allows users to create, update, delete, and list tasks based on different statuses (To-Do, In-Progress, Done).

Project Requirements: https://roadmap.sh/projects/task-tracker

---

## Features

- **Add a task** with a description.
- **Update a task's description** or status.
- **Mark a task as "In-Progress" or "Done".**
- **Delete a task.**
- **List tasks** based on status (`toDo`, `inProgress`, `done`).

---

## Usage

### Commands

You can run the commands from the terminal using the following format:

go run main.go <command> <arguments>


### Available Commands:

1. **Add Task:**

    Create a new task with a description.

    ```bash
    go run main.go add <task description>
    ```

    Example:
    ```bash
    go run main.go add "Finish Go project"
    ```

2. **Update Task:**

    Update the description of a task (or set a new description).

    ```bash
    go run main.go update <task ID> <new description>
    ```

    Example:
    ```bash
    go run main.go update 12345 "Updated task description"
    ```

3. **Mark Task as In-Progress:**

    Mark a task as "In-Progress".

    ```bash
    go run main.go mark-in-progress <task ID>
    ```

4. **Mark Task as Done:**

    Mark a task as "Done".

    ```bash
    go run main.go mark-done <task ID>
    ```

5. **Delete Task:**

    Delete an existing task.

    ```bash
    go run main.go delete <task ID>
    ```

6. **List Tasks:**

    List tasks by status (`toDo`, `inProgress`, `done`) or all tasks.

    ```bash
    go run main.go list <status (toDo | inProgress | done)>
    ```

    Example:
    ```bash
    go run main.go list toDo
    ```

    If no status is provided, all tasks will be listed.

    ```bash
    go run main.go list
    ```

---


