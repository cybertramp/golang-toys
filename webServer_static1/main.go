package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

var (
    timeSumsMu sync.RWMutex
    timeSums   int64
    port       uint16
)

func main() {
    port = 8080
    fmt.Println(port)
    
    // Start the goroutine that will sum the current time
    // once per second.
    go runDataLoop()

    // Create a handler that will read-lock the mutext and
    // write the summed time to the client
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        timeSumsMu.RLock()
        defer timeSumsMu.RUnlock()
        fmt.Fprint(w, timeSums)
    })
    http.ListenAndServe(":8080", nil)

}

func runDataLoop() {
    for {
        // Within an infinite loop, lock the mutex and
        // increment our value, then sleep for 1 second until
        // the next time we need to get a value.
        timeSumsMu.Lock()
        timeSums += time.Now().Unix()
        timeSumsMu.Unlock()
        time.Sleep(1 * time.Second)
    }
}
