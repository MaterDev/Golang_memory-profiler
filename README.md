# Go Memory Profiling Tutorial (Explained)

This project demonstrates how to use Go's built-in tools to profile memory usage in a simple HTTP server application. It's designed for beginners to understand memory concepts and profiling in Go.

## Key Concepts

- **Memory Allocation**: The process of reserving a portion of computer memory for program use.
- **Heap**: A region of a program's memory used for dynamic allocation. Unlike the stack, heap memory is not automatically freed when a function returns.
- **Garbage Collection**: An automatic memory management process that frees heap memory that's no longer in use.
- **Profiling**: The process of analyzing a program's behavior, particularly its memory usage and performance.

## Project Structure

```
go-memory-profiling/
├── main.go               # The entry point of our application
├── handler/
│   └── handler.go        # Contains HTTP request handlers
├── profiling/
│   └── profiling.go      # Sets up profiling tools
├── go.mod                # Go module file
└── README.md             # Project documentation
```

## Prerequisites

- Go 1.16 or later
- Git

## Setup

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/go-memory-profiling.git
   cd go-memory-profiling
   ```

2. Initialize the Go module:
   ```
   go mod init github.com/yourusername/go-memory-profiling
   ```

## Running the Application

To run the application:

```
go run main.go
```

The server will start on `http://localhost:8080`.

## Understanding the Code

Each file in this project contains detailed comments explaining what each part does and why. Be sure to read through the code comments to understand how memory allocation and profiling work in Go.

## Profiling

1. While the application is running, access the pprof interface at `http://localhost:8080/debug/pprof/`.

2. Generate a heap profile (shows memory allocation):
   ```
   go tool pprof http://localhost:8080/debug/pprof/heap
   ```

3. Generate a 30-second CPU profile (shows where the program spends its time):
   ```
   go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
   ```

4. View goroutines (lightweight threads managed by Go runtime):
   ```
   go tool pprof http://localhost:8080/debug/pprof/goroutine
   ```

## Analyzing Profiles

After generating a profile, you can use these commands in the pprof interactive mode:

- `top`: Shows the top memory-consuming functions
- `web`: Generates a graph visualization (requires Graphviz)
- `list <function>`: Shows line-by-line memory usage for a function
