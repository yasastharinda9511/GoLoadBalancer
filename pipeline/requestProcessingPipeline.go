package pipeline

import (
	"fmt"
	"net/http"

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
	println("Processing request...")
	ruleID, ruleErr := p.ruleStore.Evaluate(requestMessage)

	if ruleErr != nil {
		p.handleError(ruleErr, requestMessage)
		return
	}

	pool, poolErr := p.poolSelector.GetPool(ruleID)

	if poolErr != nil {
		p.handleError(poolErr, requestMessage)
		return
	}

	statusCode, body, backendErr := pool.HandleBackendCall(requestMessage)

	if backendErr != nil {
		p.handleError(backendErr, requestMessage)
		return
	}

	responseMsg := message.NewHttpResponseMessage(statusCode, body, requestMessage)

	p.ResponseProcessingPipeline.Execute(responseMsg)
}

func (p *RequestProcessingPipeline) handleError(err error, requestMessage *message.HttpRequestMessage) {
	fmt.Print("Error occurred: ")
	errorMsg := []byte(err.Error())
	responseMsg := message.NewHttpResponseMessage(http.StatusInternalServerError, errorMsg, requestMessage)
	p.ResponseProcessingPipeline.Execute(responseMsg)
}
