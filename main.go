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
	// ! Enable profiling
	profiling.EnableProfiling()

	/** 
		Setup some routes. Will use built-in api for handling server code.
		Using HandleFunc will allow for route handlers to be stored in separate files.
	*/
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/allocate", handler.AllocateHandler)
	http.HandleFunc("/allocate2", handler.AllocateHandler2)

	// * Timestamp on server start/restart to keep track of updates in console
	log.Printf("Server starting at %s", time.Now().Format(time.RFC3339))

	// Start server
	PORT := 8080
	fmt.Printf("Server is running on port: %v \n", PORT)
	fmt.Printf("Profiling Data => http://localhost:%v/debug/pprof/ \n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
	
}