# Profiling Instructions

## Key Concepts

- **Heap**: A region of a program's memory used for dynamic allocation. Objects on the heap live beyond the scope of the function that created them, unlike stack memory.

- **Allocation**: The process of requesting memory from the system. In Go, small allocations may happen on the stack, while larger or longer-lived objects are allocated on the heap.

- **Garbage Collection**: Go's automatic memory management system that periodically frees heap memory no longer in use.

- **CPU Profile**: Shows where a program spends its CPU time, useful for finding performance bottlenecks.

- **Memory Profile (Heap Profile)**: Shows memory allocation sites, useful for finding memory leaks or excessive allocations.

- **Goroutine**: A lightweight thread managed by the Go runtime. Many goroutines can run on a single OS thread.

## Project Structure

```go
package main

import (
    "fmt"
    "golang-memory-profiler/handler"
    "golang-memory-profiler/profiling"
    "log"
    "net/http"
    "time"
)

func main() {
    profiling.EnableProfiling()

    http.HandleFunc("/", handler.HelloHandler)
    http.HandleFunc("/allocate", handler.AllocateHandler)
    http.HandleFunc("/allocate2", handler.AllocateHandler2)

    log.Printf("Server starting at %s", time.Now().Format(time.RFC3339))

    PORT := 8080
    fmt.Printf("Server is running on port: %v \n", PORT)
    fmt.Printf("Profiling Data => http://localhost:%v/debug/pprof/ \n", PORT)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}
```

## Steps to Profile

1. Start the application:
   ```
   go run main.go
   ```

2. Generate load:
   - Visit `http://localhost:8080/` for the `HelloHandler`.
   - Visit `http://localhost:8080/allocate` for the `AllocateHandler`.
   - Visit `http://localhost:8080/allocate2` for the `AllocateHandler2` (direct memory measurement).

3. Capture a heap profile:
   ```
   go tool pprof http://localhost:8080/debug/pprof/heap
   ```
   In the interactive session:
   - `top`: Shows top memory-consuming functions
   - `web`: Opens a web visualization (requires Graphviz)
   - `list AllocateHandler`: Shows line-by-line memory usage

4. Capture a 30-second CPU profile:
   ```
   go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
   ```
   Generate load during this window, then analyze as above.

5. View active goroutines:
   ```
   go tool pprof http://localhost:8080/debug/pprof/goroutine
   ```

6. Memory statistics:
   Visit `http://localhost:8080/debug/pprof/` and click "heap".

## Understanding and Verifying Profile Output

1. Profile Analysis:
   - Heap profile: Identify functions allocating the most memory.
   - CPU profile: Identify functions using the most CPU time.
   - Goroutine profile: Check for potential deadlocks or goroutine leaks.

2. Direct Memory Measurement (AllocateHandler2):
   This handler provides a direct measurement of memory allocation:

   ```go
   func AllocateHandler2(w http.ResponseWriter, r *http.Request) {
       var m1, m2 runtime.MemStats
       runtime.ReadMemStats(&m1)

       data := make([]byte, 100*1024*1024)
       data[0] = 1

       runtime.ReadMemStats(&m2)

       fmt.Fprintf(w, "Heap Alloc Before: %d, After: %d, Diff: %d\n", 
           m1.HeapAlloc, m2.HeapAlloc, m2.HeapAlloc-m1.HeapAlloc)
       fmt.Fprintf(w, "Heap Sys Before: %d, After: %d, Diff: %d\n", 
           m1.HeapSys, m2.HeapSys, m2.HeapSys-m1.HeapSys)
   }
   ```

   Use this to verify actual memory allocation vs profiler reports.

3. Interpreting MemStats:
   - HeapAlloc: Actual memory allocated for heap objects.
   - HeapSys: Total heap memory obtained from the OS.
   - The difference between these values before and after allocation shows the true allocation size.

4. Profile vs Direct Measurement:
   - Compare the results from the profiler (`/allocate`) with the direct measurement (`/allocate2`).
   - If there are discrepancies, trust the direct measurement for accurate allocation sizes.

5. Continuous Monitoring:
   - The server logs its start time, which helps in tracking restarts and relating them to memory usage patterns.
   - Use the provided profiling data URL (`http://localhost:8080/debug/pprof/`) for ongoing monitoring.

Remember to test your application under realistic load conditions for the most meaningful profiling results. The combination of profiler data and direct measurements provides a comprehensive view of your application's memory behavior.