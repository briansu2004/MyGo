# Galvanize A Go REST API App

- [Step 1: Set Up Our Go Environment](#step-1-set-up-our-go-environment)
- [Step 2: Create Our Project Directory](#step-2-create-our-project-directory)
- [Step 3: Initialize Go Module](#step-3-initialize-go-module)
- [Step 4: Install Required Packages](#step-4-install-required-packages)
- [Step 5: Create Our Main File](#step-5-create-our-main-file)
- [Step 6: Write Our API Code](#step-6-write-our-api-code)
- [Step 7: Run Our Server](#step-7-run-our-server)
- [Step 8: Test Our API](#step-8-test-our-api)
  - [Step 8.1: Test With cURL](#step-81-test-with-curl)
  - [Step 8.2: Test With Postman](#step-82-test-with-postman)
  - [Step 8.3: Test With Web Browser](#step-83-test-with-web-browser)

<!-- To galvanize a Go REST API app, you'll need to set up a basic project structure, define our endpoints, handle requests and responses, and manage dependencies. Here's a simple guide to help we get started: -->

## Step 1: Set Up Our Go Environment

Ensure we have Go installed on our system. You can download and install it from the official Go website: <https://golang.org/>

## Step 2: Create Our Project Directory

Create a directory for our project. Inside this directory, we'll organize our files and folders.

```dos
cd C:\devbox-go
md go-kickstart-api-app
cd go-kickstart-api-app
```

## Step 3: Initialize Go Module

Go modules allow us to manage dependencies for our project. Run the following command in our project directory:

```dos
go mod init github.com/briansu2004/go-kickstart-api-app
```

## Step 4: Install Required Packages

For building a REST API, we need a router package. A popular choice is `gorilla/mux`. Install it using:

```dos
go get -u github.com/gorilla/mux

go install github.com/gorilla/mux@latest
```

## Step 5: Create Our Main File

Create a file named `main.go` in our project directory. This file will contain the main code for our REST API.

## Step 6: Write Our API Code

Here's a basic example of how we can set up a simple REST API using Gorilla Mux:

```go
package main

import (
"encoding/json"
"log"
"net/http"

"github.com/gorilla/mux"
)

type Item struct {
ID   string `json:"id,omitempty"`
Name string `json:"name,omitempty"`
}

var items []Item

func GetItems(w http.ResponseWriter, r *http.Request) {
json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
var newItem Item
json.NewDecoder(r.Body).Decode(&newItem)
items = append(items, newItem)
json.NewEncoder(w).Encode(newItem)
}

func main() {
router := mux.NewRouter()

// Define routes
router.HandleFunc("/items", GetItems).Methods("GET")
router.HandleFunc("/items", CreateItem).Methods("POST")

log.Fatal(http.ListenAndServe(":8000", router))
}
```

## Step 7: Run Our Server

Navigate to our project directory and run our server using:

```dos
go run main.go
```

## Step 8: Test Our API

You can use tools like cURL, Postman, or simply our web browser to test our API endpoints.

### Step 8.1: Test With cURL

Open a new terminal window and use cURL commands to send requests to your API endpoints.

```dos
// Get items
curl http://localhost:8000/items

// Create a new item (Windows)
curl -X POST -d "{\"id\":\"1\", \"name\":\"Item 1\"}" -H "Content-Type: application/json" http://localhost:8000/items
curl -X POST -d "{\"id\":\"2\", \"name\":\"Item 2\"}" -H "Content-Type: application/json" http://localhost:8000/items

// Create a new item (Unix)
curl -X POST -d '{"id":"1", "name":"Item 1"}' -H "Content-Type: application/json" http://localhost:8000/items

// Get item #1
curl http://localhost:8000/items/1
```

### Step 8.2: Test With Postman

### Step 8.3: Test With Web Browser

- <http://localhost:8000/items>
<!-- - <http://localhost:8000/items/1> -->

<!-- 
1. **Test Your Endpoints:**

You can test your API using various methods:

a. **Using cURL:** Open a new terminal window and use cURL commands to send requests to your API endpoints. For example:

```dos
// Get items
curl http://localhost:8000/items

// Create a new item (Windows)
curl -X POST -d "{\"id\":\"1\", \"name\":\"Item 1\"}" -H "Content-Type: application/json" http://localhost:8000/items

// Create a new item (Unix)
curl -X POST -d '{"id":"1", "name":"Item 1"}' -H "Content-Type: application/json" http://localhost:8000/items
```

b. **Using Postman:** If you have Postman installed, you can create requests with different HTTP methods (GET, POST, PUT, DELETE) and send them to your API endpoints (e.g., `http://localhost:8000/items`).

c. **Using Web Browser:** You can also open your web browser and visit your API endpoints. For example:

```text
http://localhost:8000/items
``` -->

<!-- 1. **Verify Responses:**

After sending requests, verify that you're receiving the expected responses. For instance, when you send a GET request to `/items`, you should receive a JSON array containing the items.

3. **Test Different Scenarios:**

Test different scenarios such as creating items, updating items, deleting items, etc., to ensure that your API behaves as expected under various conditions.

4. **Check Server Logs:**

While your server is running, you can check the logs in your terminal to see any errors or debugging information. These logs can help you troubleshoot any issues with your API. -->

<!-- By following these steps, you can run and test your Go REST API to ensure that it's functioning correctly and meeting your project requirements. -->

<!-- That's it! You've kickstarted a simple Go REST API using Gorilla Mux. From here, we can expand our API by adding more endpoints, handling different HTTP methods, integrating with databases, adding authentication, and much more, depending on our project requirements. -->
