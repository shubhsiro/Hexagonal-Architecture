// main.go
package main

import (
	"Sarva_Assignment/internal/adapter/consensus"
	"Sarva_Assignment/internal/adapter/database"
	"Sarva_Assignment/internal/adapter/logger"
	coreDB "Sarva_Assignment/internal/core/database"
	"Sarva_Assignment/internal/core/fileprocessing"
	coreLogger "Sarva_Assignment/internal/core/logger"
	"Sarva_Assignment/internal/delivery/handler"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/go-hclog"
	"log"
	"net/http"
)

func main() {
	// Initialize the logger adapter
	logs := hclog.New(&hclog.LoggerOptions{})

	// Initialize your dependencies (e.g., RAFT, Redis, logger)
	raftAdapter, err := consensus.NewRaftAdapter("node1", logs)
	if err != nil {
		log.Fatalf("Error initializing RAFT adapter: %v", err)
	}

	// Initialize the logger that deals with the logs of adapter layer
	loggerAdapter := logger.NewHCLoggerAdapter(logs)
	loggerCore := coreLogger.NewLogger(logs)

	// Initialize the Redis adapter
	redisDB := coreDB.NewRedisDB(redis.NewClient(&redis.Options{}), loggerCore)
	redisAdapter := database.NewRedisAdapter(redisDB, loggerAdapter)

	// Initialize core components
	fileProcessor := &fileprocessing.FileProcessor{}

	// Initialize HTTP handler
	fileHandler := handler.NewFileHandler(fileProcessor, raftAdapter, redisAdapter, loggerAdapter)

	// Start your application
	// Implement your HTTP server initialization logic here
	http.HandleFunc("/upload", fileHandler.UploadFile)

	// Start the HTTP server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}

}
