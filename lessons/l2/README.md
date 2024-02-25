# Go Basics: Deep Dive into Variables, Control Structures, and Functions

Welcome back to our Go programming journey! In Lesson 2, we're going to take a closer look at some of the fundamental concepts that form the backbone of Go programming: variables, control structures, functions, and Go's type system. By the end of this lesson, you'll be more familiar with how Go handles data, makes decisions, and executes code blocks. Let's dive in!

## Variables, Types, and Declarations

Go is a statically typed language, which means that variables are explicitly typed at compile time. Understanding how to declare and use variables is crucial. Here's a quick rundown:

### Integers and Floats

```go
var age int = 30
var temperature float64 = 20.5
```

In these lines, we declare an integer variable `age` and a floating-point variable `temperature`. Go supports several numeric types, but these are the basics.

## Booleans and Strings

```go
var isGoFun bool = true
var name string = "Go Programmer"
```

Here, we declare a boolean variable `isGoFun` and a string variable `name`. These types are used to handle logical conditions and textual data, respectively.

## Inferred Typing with `:=`

```go
city := "Berlin"
```

Go also supports type inference with the `:=` syntax, allowing you to declare a variable without explicitly stating its type.

## Control Structures

Control structures direct the flow of your Go programs. Let's see how Go handles common control structures.

### If-Else Statements

```go
if num%2 == 0 {
    fmt.Println(num, "is even")
} else {
    fmt.Println(num, "is odd")
}
```

This `if-else` statement checks if a number is even or odd. Go's syntax for these statements is straightforward and similar to other C-like languages.

### For Loops

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

Go simplifies looping constructs by only providing the `for` loop, which can be adapted for a variety of use cases, including traditional counted loops as shown above.

### Switch Statements

```go
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
// Add more cases as needed
default:
    fmt.Println("Other day")
}
```

`Switch` statements in Go are used for multiple condition checks. They're cleaner and more readable, especially when comparing a single variable against multiple values.

## Functions in Go

Functions are building blocks in Go. They allow you to encapsulate reusable code blocks. Here's how to define and use them:

### Defining a Simple Function

```go
func add(x int, y int) int {
    return x + y
}
```

This function, `add`, takes two integers as input and returns their sum. Notice how Go requires you to specify the type of each parameter and the return type of the function.

### Multiple Return Values

```go
func swap(x, y string) (string, string) {
    return y, x
}
```

One of Go's unique features is support for multiple return values from a single function. This can be particularly useful for returning a result and an error value from a function.

## Running Lesson 2

To see all these concepts in action, run the lesson with the following command:

```bash
go run lessons/l2/lesson2.go
```

You'll see the output from the variable declarations, control structures, and function calls. Experiment with the code to see how changes affect the program's output.

## Conclusion

Congratulations on completing Lesson 2 of our Go series! You've now got a solid understanding of some of Go's core concepts, which will serve as the foundation for more advanced topics in the future. Stay tuned for more lessons, and happy coding!

Remember, the best way to learn is by doing. Don't hesitate to modify the examples and experiment on your own. Go's simplicity and power are in your hands now. Enjoy the journey!