// lesson2.go
package main

import (
    "fmt"
)

// Demonstrating variables, types, and declarations
func variablesAndTypes() {
    // Integer
    var age int = 30
    fmt.Println("Age:", age)

    // Floating point
    var temperature float64 = 20.5
    fmt.Println("Temperature:", temperature)

    // Boolean
    var isGoFun bool = true
    fmt.Println("Is Go fun?", isGoFun)

    // String
    var name string = "Go Programmer"
    fmt.Println("Name:", name)

    // Inferred type with :=
    city := "Berlin"
    fmt.Println("City:", city)
}

// Demonstrating control structures: if, else, for, switch
func controlStructures() {
    // if-else
    num := 10
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // for loop
    fmt.Println("Counting to 5:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // switch
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    // Add more cases as needed
    default:
        fmt.Println("Other day")
    }
}

// Demonstrating functions
func add(x int, y int) int {
    return x + y
}

// Demonstrating multiple return values
func swap(x, y string) (string, string) {
    return y, x
}

// Main function to call the demonstrations
func main() {
    fmt.Println("Lesson 2: Go Basics")

    fmt.Println("\n--- Variables and Types ---")
    variablesAndTypes()

    fmt.Println("\n--- Control Structures ---")
    controlStructures()

    fmt.Println("\n--- Functions ---")
    result := add(5, 7)
    fmt.Println("5 + 7 =", result)

    fmt.Println("\n--- Multiple Return Values ---")
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
