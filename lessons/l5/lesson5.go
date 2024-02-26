package main

import (
    "bufio"
    "fmt"
	"net/http"
	"encoding/json"
	"math"
	"strings"
    "io/ioutil"
    "os"
)

func workWithFiles() {
    // Writing to a file
    content := []byte("Hello, Go!\n")
    err := ioutil.WriteFile("example.txt", content, 0644)
    if err != nil {
        panic(err)
    }

    // Reading from a file
    data, err := ioutil.ReadFile("example.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println("File content:", string(data))

    // Using bufio to read a file line by line
    file, err := os.Open("example.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func startHTTPServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to the Go HTTP server!")
    })

    fmt.Println("Starting server at port 8080")
    http.ListenAndServe(":8080", nil)
}

func makeHTTPRequest() {
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println("Server response:", string(body))
}

type Message struct {
    Text string `json:"text"`
}

func encodeDecodeJSON() {
    // Encoding JSON
    msg := Message{Text: "Hello, JSON in Go!"}
    jsonBytes, err := json.Marshal(msg)
    if err != nil {
        panic(err)
    }
    fmt.Println("JSON:", string(jsonBytes))

    // Decoding JSON
    var decodedMsg Message
    err = json.Unmarshal(jsonBytes, &decodedMsg)
    if err != nil {
        panic(err)
    }
    fmt.Println("Decoded message:", decodedMsg.Text)
}

func useCommonPackages() {
    fmt.Println("Math Pi:", math.Pi)
    fmt.Println("String Title:", strings.Title("go is awesome"))
}

func main() {
    fmt.Println("Working with Files and I/O")
    workWithFiles()

    // Uncomment the server start in a real application
    // go startHTTPServer()
    // makeHTTPRequest()

    fmt.Println("\nEncoding/Decoding JSON")
    encodeDecodeJSON()

    fmt.Println("\nUsing Common Packages")
    useCommonPackages()
}
