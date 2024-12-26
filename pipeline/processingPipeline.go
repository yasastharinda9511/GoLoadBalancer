package pipeline

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/yasastharinda9511/go_gateway_api/dispatcher"
	"github.com/yasastharinda9511/go_gateway_api/message"
	"github.com/yasastharinda9511/go_gateway_api/pool"
	"github.com/yasastharinda9511/go_gateway_api/ruleStore"
)

// ProcessingPipeline defines the structure for the processing pipeline
type ProcessingPipeline struct {
	ruleStore    *ruleStore.RuleStore
	poolSelector *pool.PoolSelector
}

// NewProcessingPipeline creates a new instance of ProcessingPipeline
func NewProcessingPipeline(ruleStore *ruleStore.RuleStore, poolSelector *pool.PoolSelector) *ProcessingPipeline {
	return &ProcessingPipeline{
		ruleStore:    ruleStore,
		poolSelector: poolSelector,
	}
}

// Execute processes the HTTP request
func (p *ProcessingPipeline) Execute(requestMessage *message.HttpRequestMessage) {
	// Add your processing logic here
	rw := requestMessage.GetResponseWriter()
	println("Processing request...")
	dispatch := dispatcher.NewHTTPDispatcher(10 * time.Second)
	ruleID := p.ruleStore.Evaluate(requestMessage)
	pool, _ := p.poolSelector.GetPool(ruleID)

	if pool != nil {
		backend, _ := pool.Next()
		println("Backend: ", backend.GetURL())

		// Make the backend call
		resp, err := dispatch.CallBackend(dispatcher.GET, backend.GetURL(), requestMessage.GetHeaders(), requestMessage.GetQueryParams())
		if err != nil {
			// Write an error response
			http.Error(rw, "Failed to call backend", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read backend response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(rw, "Failed to read backend response", http.StatusInternalServerError)
			return
		}

		// Write the backend response to the client
		rw.Header().Set("Content-Type", "application/json") // Adjust as per backend response type
		rw.WriteHeader(http.StatusOK)
		_, writeErr := rw.Write(body)
		if writeErr != nil {
			// Log the error
			fmt.Printf("Error writing response: %v\n", writeErr)
		}

		println("Response received from backend and written to client")
	}

}
