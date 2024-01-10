// consensus.go
package consensus

import (
	"Sarva_Assignment/internal/core/logger"
	"strings"
	"sync"
	"time"
)

// Consensus represents the core logic for consensus interactions.
type Consensus struct {
	raftMessages []string
	mu           sync.Mutex
	logger       *logger.Logger // Assuming your logger is named Logger
}

// NewConsensus creates a new instance of Consensus.
func NewConsensus(logger *logger.Logger) *Consensus {
	return &Consensus{
		raftMessages: make([]string, 0),
		logger:       logger,
	}
}

// SendMessage sends a message for consensus.
func (c *Consensus) SendMessage(message string) error {
	// Placeholder: Implement your core logic for sending messages to achieve consensus
	c.logger.Log("Sending message for consensus: " + message)

	// Simulate consensus process by adding the message to the list
	c.mu.Lock()
	c.raftMessages = append(c.raftMessages, message)
	c.mu.Unlock()

	c.logger.Log("Message sent successfully.")
	return nil
}

// RetrieveMessages retrieves consensus messages.
func (c *Consensus) RetrieveMessages() ([]string, error) {
	// Placeholder: Implement your core logic for retrieving consensus messages
	c.logger.Log("Retrieving consensus messages")

	// Simulate consensus process by returning the stored messages
	c.mu.Lock()
	defer c.mu.Unlock()
	messages := make([]string, len(c.raftMessages))
	copy(messages, c.raftMessages)

	c.logger.Log("Retrieved messages: " + strings.Join(messages, ", "))
	return messages, nil
}

// StartCluster starts the consensus cluster.
func (c *Consensus) StartCluster() {
	// Placeholder: Implement your logic to start the consensus cluster
	c.logger.Log("Starting the consensus cluster")

	// Simulate some processing time
	time.Sleep(1 * time.Second)

	c.logger.Log("Consensus cluster started successfully.")
}

// Shutdown shuts down the consensus cluster.
func (c *Consensus) Shutdown() {
	// Placeholder: Implement your logic to gracefully shut down the consensus cluster
	c.logger.Log("Shutting down the consensus cluster")

	// Simulate some processing time
	time.Sleep(1 * time.Second)

	c.logger.Log("Consensus cluster shut down successfully.")
}
