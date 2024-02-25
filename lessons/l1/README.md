# Welcome to Go: Your First Step into Go Programming

Welcome to the exciting world of Go programming! If you're reading this, you're likely ready to dive into your first Go lesson. Go, or Golang, is renowned for its simplicity, efficiency, and strong support for concurrent programming. Today, we're starting with the basics: your very first Go program.

This tutorial will guide you through the content of `lesson1.go`, a simple program that introduces you to Go's syntax, variable declarations, and control structures. By the end of this lesson, you'll have a solid foundation and a running Go program that says "Hello, World!", counts numbers, and declares variables in different ways.

## Setting the Stage: Hello, World!

Every journey into a new programming language starts with a greeting. The "Hello, World!" program is a tradition, and in Go, it's as simple and elegant as the language itself:

```go
fmt.Println("Lesson 1: Hello, World!")
```

The `fmt` package is Go's way of handling formatted I/O (Input/Output). The `Println` function from the `fmt` package prints its arguments to the console, followed by a newline.

## Declaring Variables in Go

Go is statically typed, which means every variable has a type that is known at compile time. However, Go offers flexibility in variable declaration and initialization. Let's look at two ways to declare variables:

The Standard Way

```go
var message string = "Go is fun!"
fmt.Println(message)
```

Here, we explicitly state the type of the variable `message` as `string`. This method is clear and expressive, especially when you're starting out or when the type of the variable isn't obvious from the context.

The Shorthand Method

```go
shorthandMessage := "Learning Go is exciting!"
fmt.Println(shorthandMessage)
```

For a more concise declaration, Go allows you to use `:=`, which infers the type of the variable from the right-hand side of the assignment. This shorthand method is idiomatic Go and keeps your code clean and readable.

## Looping in Go: A Simple For Loop

Unlike many languages that offer a variety of looping constructs (`for`, `while`, `do-while`), Go simplifies this down to one powerful "Swiss Army knife" - the `for` loop.

```go
fmt.Println("Counting from 1 to 5:")
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

This loop initializes a counter (`i := 1`), checks a condition (`i <= 5`), and increments the counter (`i++`) on each iteration, printing the counter's value. It showcases Go's approach: doing more with less.

## Running Your First Go Program

To run this program, navigate to your project root in the terminal and execute:

```bash
go run lessons/l1/lesson1.go
```

You'll see the outputs we've discussed: a greeting to the world, two messages declared in different ways, and a countdown from 1 to 5.

Wrapping Up
Congratulations! You've just written and run your first Go program. This lesson was your first step into Go, demonstrating its straightforward syntax, the process of declaring variables, and the use of control structures like the `for` loop.

Go's simplicity doesn't stop here. As you progress, you'll discover more about its efficient concurrency model, robust standard library, and powerful interfaces. Stay tuned for more lessons as we dive deeper into Go, building up from these foundational concepts to more complex applications.