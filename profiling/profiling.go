package profiling

import (
	//  Underscore in front of an import means Golang compiler will import the packages eventhough none of its exports are directly used in our code. Any init() function in the import will still be called, which is necessary for profiling parts of the system.
	// ! This registers various http endpoint handlers for different types of profiling data.
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

/*
	Sets up profiling endpoints and configuration
		* Profiling is a technique to analyze the performance of an application.
		It can help identify bottlenecks, memory leaks, and other issues.
*/
func EnableProfiling() {
	/*
		MemProfileRate controls what fraction of memory allocations are being recorded.
			The profiler aims to sample an average of all allocation per MemProfileRate bytes allocated.

		Setting it to 1 means all allocations are recorded.
			In a production environment, this should be a higher value. (examples below)
	*/
	
	// Profile every allocation
	runtime.MemProfileRate = 1

	 // Set to profile (on average) one allocation per 1024 bytes
	//  runtime.MemProfileRate = 1024

	// Set to profile (on average) one allocation per 1 MB
    // runtime.MemProfileRate = 1024 * 1024
}

	// The pprof handlers are automatically registered via the blank import of "net/http/pprof" above.
	// This adds several endpoints under /debug/pprof/, including:
	// - /debug/pprof/heap: A sampling of memory allocations
	// - /debug/pprof/goroutine: Stack traces of all current goroutines
	// - /debug/pprof/profile: CPU profile
	// - /debug/pprof/block: Stack traces that led to blocking on synchronization primitives
	// - /debug/pprof/threadcreate: Stack traces that led to the creation of new OS threads

// Add this function to explicitly write a heap profile
func WriteHeapProfile(filename string) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    runtime.GC() // Force garbage collection to get up-to-date statistics
    if err := pprof.WriteHeapProfile(f); err != nil {
        return err
    }
    return nil
}