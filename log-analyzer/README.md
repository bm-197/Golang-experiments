# Log Analyzer

This project was created to practice concurrency in Go. It processes and aggregates metrics from HTTP logs efficiently using Go's concurrency features.

## Features

- **Concurrent Log Processing**: Analyzes multiple log files in parallel using goroutines.
- **Aggregated Metrics**: Reports total request count and average response time per endpoint.
- **Command-Line Input**: Accepts multiple log files via command-line arguments.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) must be installed.

### Steps

1. **Clone the repository**:
   ```bash
   git clone https://github.com/bm-197/log-analyzer.git
   ```
2. **Navigate to the project directory**:
    ```bash
    cd log-analyzer
    ```
3. **Build the project**:
    ```bash
    go build -o log-analyzer
    ```
4. **Run the log analyzer**:
    ```bash
    ./log-analyzer <log_file1> <log_file2> ... <log_fileN>
    ```
    Example:

    ```bash
    ./log-analyzer server_log1.txt server_log2.txt
    ```

## Usage

To use the log analyzer, specify one or more log files as command-line arguments. The program will process the logs and print a summary of the following metrics for each endpoint:

- Total number of requests
- Average response time (in milliseconds)

### Example

If you have three log files (`server_log1.txt`, `server_log2.txt`, and `server_log3.txt`), run the tool like this:

```bash
./log-analyzer server_log1.txt server_log2.txt server_log3.txt
```
### Log File Format

The log files should have each entry in the following format:

    ```bash
    <HTTP_METHOD> <ENDPOINT> <RESPONSE_TIME_IN_MS>
    ```

### Example log entries

    ```bash
    GET /api/users 120
    POST /api/products 400
    DELETE /api/users/1 250
    ```

## Explanation

- **Goroutines**: The tool processes each log file concurrently using goroutines to improve performance.
- **Channels**: Aggregated metrics from each goroutine are sent back to the main thread via channels to avoid race conditions.
- **Command-line Arguments**: The program accepts any number of log files as input, making it flexible for various use cases.



