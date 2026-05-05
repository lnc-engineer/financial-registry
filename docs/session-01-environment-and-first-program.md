concept

A Go project start with a module (go.mod) which defines the project

The main function is the entry point of a program
When running a Go program, execution starts from main()

How it works

Creation of a Go module using:

go mod init github.com/lnc-engineer/financial-registry

This sets the identity of the project

The structure

cmd/hello/main.go

cmd/ is used for executable programs

Running:

go run ./cmd/hello

compiles and runs the program 

TEST

Ran the program using:

go run ./cmd/hello

Output:
"Financial Systems Registry - Initialised"

Confirmed:
-Go is installed correctly
-Project structure works
-Program executes from main()


Code pattern

package main

import "fmt"

func main() {
    fmt.Println("Financial Systems Registry - Initialised")
}



System connection

This is the startng point of the Financial Systems Registry

It represents the simplest executable component of a system

All future services (ingestion, APIs, workder) will follow this structure



