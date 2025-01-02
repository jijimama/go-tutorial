package main

import "fmt"

import "go-tutorial/greetings"

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
