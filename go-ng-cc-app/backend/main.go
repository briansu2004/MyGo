package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/rs/cors"
)

// Global DB connection
var db *sql.DB

// Connection String (Update with your credentials!)
// const connString = "server=localhost;user id=sa;password=YourStrongPassword123;database=GoAppDB"
const connString = "server=localhost;port=51109;database=GoAppDB;trusted_connection=yes;TrustServerCertificate=true"

// Job represents the work to be done
type Job struct {
	ID   int
	Name string
}

// Worker function: Runs inside a Goroutine
func worker(workerID int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup this worker is done when function exits

	for job := range jobs {
		// Simulate expensive processing (e.g., complex calculation)
		time.Sleep(500 * time.Millisecond)

		fmt.Printf("Worker %d started job %s\n", workerID, job.Name)

		// Insert into MS SQL Server
		query := "INSERT INTO ProcessedJobs (JobName, ProcessedByWorkerID, ProcessedAt) VALUES (@p1, @p2, GETDATE())"
		_, err := db.Exec(query, job.Name, workerID)
		if err != nil {
			log.Println("Error inserting to DB:", err)
		}
	}
}

// // Handler to trigger the batch process
// func processBatchHandler(w http.ResponseWriter, r *http.Request) {
// 	startTime := time.Now()

// 	const numJobs = 20
// 	const numWorkers = 5 // We will use 5 concurrent threads

// 	// 1. Create Channels
// 	// Buffered channel to hold the jobs
// 	jobsChannel := make(chan Job, numJobs)

// 	// 2. Setup WaitGroup to track completion
// 	var wg sync.WaitGroup

// 	// 3. Start Workers (Goroutines)
// 	// This spins up 5 separate threads waiting for work
// 	for w := 1; w <= numWorkers; w++ {
// 		wg.Add(1)
// 		go worker(w, jobsChannel, &wg)
// 	}

// 	// 4. Send work to the jobs channel
// 	for i := 1; i <= numJobs; i++ {
// 		jobsChannel <- Job{ID: i, Name: fmt.Sprintf("Order_#%d", i)}
// 	}

// 	// 5. Close channel so workers know no more jobs are coming
// 	close(jobsChannel)

// 	// 6. Wait for all workers to finish
// 	wg.Wait()

// 	elapsed := time.Since(startTime)
// 	response := map[string]string{
// 		"message": fmt.Sprintf("Processed %d jobs using %d workers", numJobs, numWorkers),
// 		"time":    elapsed.String(),
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func processBatchHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for Server-Sent Events (SSE)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, _ := w.(http.Flusher)

	const numJobs = 20
	const numWorkers = 5
	jobsChannel := make(chan Job, numJobs)
	resultsChannel := make(chan string, numJobs) // New channel for progress updates
	var wg sync.WaitGroup

	// Start Workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobsChannel {
				time.Sleep(500 * time.Millisecond) // Simulate work

				// Perform DB Insert
				db.Exec("INSERT INTO ProcessedJobs (JobName, ProcessedByWorkerID) VALUES (@p1, @p2)", job.Name, id)

				// Send a message back to the results channel
				resultsChannel <- fmt.Sprintf("Worker %d finished %s", id, job.Name)
			}
		}(w)
	}

	// Feed the jobs
	for i := 1; i <= numJobs; i++ {
		jobsChannel <- Job{ID: i, Name: fmt.Sprintf("Order_#%d", i)}
	}
	close(jobsChannel)

	// Monitor progress in a separate goroutine
	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	// Read from resultsChannel and PUSH to the UI immediately
	for msg := range resultsChannel {
		fmt.Fprintf(w, "data: %s\n\n", msg) // SSE format
		flusher.Flush()                     // Send it to the browser NOW
	}
}

func main() {
	var err error
	// Connect to Database
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}
	fmt.Println("Connected to MS SQL Server!")

	// Setup Router
	mux := http.NewServeMux()
	mux.HandleFunc("/process", processBatchHandler)

	// Setup CORS (Allow Angular localhost:4200)
	handler := cors.Default().Handler(mux)

	fmt.Println("Server starting on port 8091...")
	log.Fatal(http.ListenAndServe(":8091", handler))
}
