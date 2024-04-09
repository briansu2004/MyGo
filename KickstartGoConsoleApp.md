# Kickstart a Go console application

Kickstarting a Go console application involves setting up the project, defining the application logic, and handling command-line arguments.

Here's a step-by-step guide for an example - the program checks if there's at least one command-line argument. If there is, it prints a greeting message with the provided name. Otherwise, it displays usage instructions.

## Step 1: Setup Go environment

Make sure We have Go installed on our system. We can download it from the [official Go website](https://golang.org/dl/) and follow the installation instructions.

## Step 2: Create a new Go module

Go modules manage dependencies for our project. Create a new directory for our project and initialize it as a Go module using the following command:

```dos
cd C:\devbox-go
md go-kickstart-console-app
cd go-kickstart-console-app
go mod init github.com/briansu2004/go-kickstart-console-app
```

## Step 3: Create the main file

Create a file named `main.go` within our project directory. This will be the entry point for our application.

## Step 4: Define our application logic

Write our application logic in the `main.go` file. This may include initializing variables, defining functions, and setting up any necessary data structures.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Access command-line arguments
    args := os.Args

    // Check if there are any arguments
    if len(args) < 2 {
        fmt.Println("Usage: myconsoleapp [name]")
        return
    }

    // Extract the first argument (excluding the program name)
    name := args[1]

    // Print a greeting message
    fmt.Printf("Hello, %s!\n", name)
}
```

## Step 5: Handle command-line arguments

Use the `os.Args` variable to access command-line arguments passed to our application. We can parse these arguments using the `flag` package or manually.

## Step 6: Build and run our application

After writing our code, build our application using the `go build` command:

```dos
go build
```

This will create an executable file in our project directory. We can then run our application:

```dos
go-kickstart-console-app.exe
```

```text
C:\devbox-go\go-kickstart-console-app>go-kickstart-console-app Brian Su
Hello, Brian!
```
