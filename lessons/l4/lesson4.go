// lesson4.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Demonstrating basic goroutines
func sayHello(i int, wg *sync.WaitGroup) {
    defer wg.Done() // Notify the WaitGroup that we're done
    time.Sleep(100 * time.Millisecond) // Simulate work
    fmt.Printf("Hello from goroutine %d\n", i)
}

// Demonstrating channels
func count(thing string, c chan string) {
    for i := 1; i <= 5; i++ {
        c <- fmt.Sprintf("%s %d", thing, i) // Send a value into the channel
        time.Sleep(time.Millisecond * 500)  // Simulate work
    }
    close(c) // Close the channel
}

// Demonstrating select statement for non-blocking channel operations
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("Quit")
            return
        }
    }
}

func main() {
    var wg sync.WaitGroup

    fmt.Println("--- Goroutines ---")
    for i := 1; i <= 5; i++ {
        wg.Add(1) // Increment the WaitGroup counter
        go sayHello(i, &wg)
    }
    wg.Wait() // Wait for all goroutines to finish

    fmt.Println("\n--- Channels ---")
    things := make(chan string)
    go count("sheep", things)
    for thing := range things { // Receive values from the channel
        fmt.Println(thing)
    }

    fmt.Println("\n--- Select Statement ---")
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
