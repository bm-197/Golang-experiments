# CLI To-Do App

This project was created to practice structs, methods, functions, and data types in Go. It is a simple command-line to-do app that allows you to manage your tasks directly from the terminal using various commands.

## Features

- **Add Tasks**: Easily add new tasks to your to-do list.
- **Edit Tasks**: Modify the title of an existing task using its index.
- **Delete Tasks**: Remove tasks by specifying their index.
- **Toggle Task Completion**: Mark a task as complete or incomplete by toggling its state.
- **List Tasks**: Display all tasks in your list.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) must be installed on your machine.

### Steps

1. **Clone the repository**:
   ```bash
   git clone https://github.com/bm-197/Golang-experiments
   ```
2. **Navigate to the project directory:**
    ```bash
    cd cli-todo-app
    ```
3. **Build the project**
    ```bash
    go build -o todo
    ```
4. **Run the CLI To-Do App:**
    ```bash
    ./todo [flags]
    ```
### Usage
    Use the following commands to manange your to-do list
    - **Add a task:**
        ```bash
        ./todo -add "New Task"
        ```
    - **Edit a task:**
        ```bash
        ./todo -edit 1:"Updated Task Title"
        ```
    - **Delete a task by index:**
        ```bash
        ./todo -del 2
        ```
    - **Toggle a task by index (complete/incomplete):**
        ```bash
        ./todo -toggle 1
        ```
    - **List all tasks:**
        ```bash
        ./todo -list
        ```

