# Kickstart a Go console application

- [Step 1: Set Up Our Workspace](#step-1-set-up-our-workspace)
- [Step 2: Initialize Go Modules (Optional but Recommended)](#step-2-initialize-go-modules-optional-but-recommended)
- [Step 3: Install Dependencies](#step-3-install-dependencies)
- [Step 4: Write Our Application Code](#step-4-write-our-application-code)
- [Step 5: Run Our Application](#step-5-run-our-application)
- [Step 6: Test Our Application](#step-6-test-our-application)
- [Step 7: Deploy Our Application (Optional)](#step-7-deploy-our-application-optional)

Kickstarting a Go web application involves setting up the basic structure, handling dependencies, defining routes, and starting the server.

## Step 1: Set Up Our Workspace

Create a new directory for our project. Inside this directory, create a file named `main.go` which will contain our application code.

## Step 2: Initialize Go Modules (Optional but Recommended)

<!-- If we are using Go version 1.11 or later, we can initialize Go modules by running: -->

```dos
cd C:\devbox-go
md go-kickstart-web-app
cd go-kickstart-web-app
go mod init github.com/briansu2004/go-kickstart-web-app
```

This will create a `go.mod` file in our project directory to manage dependencies.

## Step 3: Install Dependencies

We need a third-party library - Gorilla Mux router - for our web application. Use `go get` to install dependencies. 

```dos
go get github.com/gorilla/mux
```

## Step 4: Write Our Application Code

Create a new file `main.go`

<!-- , import necessary packages and define our application logic. -->

<!-- Here's a basic example using the Gorilla Mux router: -->

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)

    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to my Go web app!")
}
```

## Step 5: Run Our Application

To run our Go web application, simply execute:

```dos
go run main.go
```

We should see the message "Server is running on <http://localhost:8080>".

## Step 6: Test Our Application

Open our web browser and go to `http://localhost:8080` to see our web application in action.

## Step 7: Deploy Our Application (Optional)

Once our web application is ready, we might want to deploy it to a server. We can use various methods like Docker, traditional hosting, or cloud platforms like AWS, Google Cloud, or Heroku.

<!-- This is a basic example to get we started. As our application grows, we might want to add more features, middleware, database integration, and so on. Go has a rich ecosystem and plenty of resources available online to help we build robust web applications. -->
