// lesson1.go
package main

import (
    "fmt"
)

func main() {
    // Print "Hello, World!" to the console
    fmt.Println("Lesson 1: Hello, World!")

    // Declare a variable
    var message string = "Go is fun!"
    fmt.Println(message)

    // Using shorthand variable declaration
    shorthandMessage := "Learning Go is exciting!"
    fmt.Println(shorthandMessage)

    // Simple for loop in Go, printing numbers 1 through 5
    fmt.Println("Counting from 1 to 5:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}
