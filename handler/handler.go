package handler

import (
	"fmt"
	"net/http"
)

/* 
	HelloHandler() will respond with a simple Hello World, to demonstrate server is working
		An initial health check that things are up and running.
*/
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	// * Logic to specify this route will only work for GET request methods.
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed on '/'", http.StatusMethodNotAllowed)

		return
	}

	fmt.Fprintf(w, "Hello World 🌎 - From Server")
}

/*
	AllocateHandler() will demonstrate memory allocation.
		This function allocates a large chunk of memory to show how it appears in memory profiles.
*/
func AllocateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed on '/allocate'", http.StatusMethodNotAllowed)

		return
	}

	fmt.Fprintf(w, "Allocate Route activated. 💾 \n")

	/*
		Allocate 100 MB Slice
			A slice of bytes with a length of 100MB. This will be a continuous block of memory.
			This memory is allocated on the heap because its too large for the stack.

			THe make() function allocates memory and also intializes it with zero-valuies.
				* The zero-value of a byte is an 8bit version of 0, which is: 00000000). 

				? make([]byte, 100*1024*1024)
					1024 is the number of bytes in a kilobyte (KB)
					1024*1024 is the number of bytes in a megabyte (MB)
					100*1024*1024 is 100 megabytes
				
					🧠 1024*1024*1024 would be 1 gigabyte (GB)

			Heap: Region of memory used for dynamic allocation.
			Stack: Region of memory used for static allocation (like fuinction calls and local variables)
	*/
	data := make([]byte, 100*1024*1024)

	// This is just doing somethign with the data to prevent the Go compiler from optimizing it away.
	data[0] = 1

	// Respond to HTTP request
	fmt.Fprintf(w, "Allocated 100 MB \n")

	// Once the slice goes out of scope, such as when the function returns, it become eligible for garbage collection. This may not happen immediately.

	/*
		Such large slices might be used for 
			- Buffering large files
			- Image processing
			- Temporary storage for large data sets.
	*/
}
