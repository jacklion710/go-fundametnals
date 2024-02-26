// lesson3.go
package main

import (
    "fmt"
)

// Demonstrating arrays and slices
func arraysAndSlices() {
    // Array declaration
    var numbers [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array:", numbers)

    // Slices are a more flexible and powerful interface to sequences in Go.
    var slice []int = numbers[1:4] // Creates a slice from numbers[1] to numbers[3]
    fmt.Println("Slice:", slice)

    // Append to a slice
    slice = append(slice, 10)
    fmt.Println("Slice after append:", slice)
}

// Demonstrating maps
func mapsDemo() {
    // Map declaration
    var colors map[string]string = map[string]string{
        "red":   "#ff0000",
        "green": "#00ff00",
        "blue":  "#0000ff",
    }
    fmt.Println("Map:", colors)

    // Adding a new key-value pair
    colors["white"] = "#ffffff"
    fmt.Println("Map after adding 'white':", colors)

    // Deleting a key-value pair
    delete(colors, "green")
    fmt.Println("Map after deleting 'green':", colors)
}

// Demonstrating pointers
func pointersDemo() {
    x := 10
    var p *int = &x // p holds the address of x
    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", p)
    fmt.Println("Value at address p:", *p)

    *p = 20 // Changing the value at the address p points to
    fmt.Println("New value of x:", x)
}

// Demonstrating type embedding and composition
type Engine struct {
    Power int
    Type  string
}

type Car struct {
    Engine // Embedded type
    Make   string
    Model  string
}

func typeEmbeddingAndComposition() {
    myCar := Car{
        Engine: Engine{
            Power: 150,
            Type:  "Petrol",
        },
        Make:  "Toyota",
        Model: "Corolla",
    }

    fmt.Println("My car:", myCar)
    fmt.Println("Engine Power:", myCar.Power) // Accessing Engine's Power field directly through Car
}

// Main function to call all demonstrations
func main() {
    fmt.Println("Lesson 3: Advanced Types")

    fmt.Println("\n--- Arrays and Slices ---")
    arraysAndSlices()

    fmt.Println("\n--- Maps ---")
    mapsDemo()

    fmt.Println("\n--- Pointers ---")
    pointersDemo()

    fmt.Println("\n--- Type Embedding and Composition ---")
    typeEmbeddingAndComposition()
}
