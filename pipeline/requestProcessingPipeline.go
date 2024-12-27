package pipeline

import (
	"io"
	"net/http"
	"time"

	"github.com/yasastharinda9511/go_gateway_api/dispatcher"
	"github.com/yasastharinda9511/go_gateway_api/message"
	"github.com/yasastharinda9511/go_gateway_api/pool"
	"github.com/yasastharinda9511/go_gateway_api/ruleStore"
)

// ProcessingPipeline defines the structure for the processing pipeline
type RequestProcessingPipeline struct {
	ruleStore                  *ruleStore.RuleStore
	poolSelector               *pool.PoolSelector
	ResponseProcessingPipeline *ResponseProcessingPipeline
}

// NewProcessingPipeline creates a new instance of ProcessingPipeline
func NewRequestProcessingPipeline(ruleStore *ruleStore.RuleStore, poolSelector *pool.PoolSelector, responseProcessinPipeline *ResponseProcessingPipeline) *RequestProcessingPipeline {
	return &RequestProcessingPipeline{
		ruleStore:    ruleStore,
		poolSelector: poolSelector,
	}
}

// Execute processes the HTTP request
func (p *RequestProcessingPipeline) Execute(requestMessage *message.HttpRequestMessage) {
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
		statusCode := resp.StatusCode
		responseMsg := message.NewHttpResponseMessage(statusCode, body, requestMessage)

		// Process the response
		p.ResponseProcessingPipeline.Execute(responseMsg)
	}

}
