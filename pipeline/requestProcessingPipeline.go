package pipeline

import (
	"fmt"

	"github.com/yasastharinda9511/go_gateway_api/message"
	"github.com/yasastharinda9511/go_gateway_api/pool"
	"github.com/yasastharinda9511/go_gateway_api/ruleStore"
)

// ProcessingPipeline defines the structure for the processing pipeline
type RequestProcessingPipeline struct {
	ruleStore                  *ruleStore.RuleStore
	poolSelector               *pool.PoolSelector
	responseProcessingPipeline *ResponseProcessingPipeline
	requestMessagePool         *message.Pool[*message.HttpRequestMessage]
	responseMessagePool        *message.Pool[*message.HttpResponseMessage]
}

// NewProcessingPipeline creates a new instance of ProcessingPipeline
func NewRequestProcessingPipeline(ruleStore *ruleStore.RuleStore,
	poolSelector *pool.PoolSelector,
	responseProcessinPipeline *ResponseProcessingPipeline,
	requestMessagePool *message.Pool[*message.HttpRequestMessage],
	responseMessagePool *message.Pool[*message.HttpResponseMessage],
) *RequestProcessingPipeline {
	return &RequestProcessingPipeline{
		ruleStore:                  ruleStore,
		poolSelector:               poolSelector,
		responseProcessingPipeline: responseProcessinPipeline,
		requestMessagePool:         requestMessagePool,
		responseMessagePool:        responseMessagePool,
	}
}

// Execute processes the HTTP request
func (p *RequestProcessingPipeline) Execute(requestMessage *message.HttpRequestMessage) {
	// Add your processing logic heredd
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

	responseMsg := p.responseMessagePool.Get()
	responseMsg.SetHttpRequestMessage(requestMessage)
	responseMsg.SetStatusCode(statusCode)
	responseMsg.SetBody(body)

	p.responseProcessingPipeline.Execute(responseMsg)
}

func (p *RequestProcessingPipeline) handleError(err error, requestMessage *message.HttpRequestMessage) {
	fmt.Print("Error occurred: ")
	errorMsg := []byte(err.Error())
	responseMsg := p.responseMessagePool.Get()
	responseMsg.SetHttpRequestMessage(requestMessage)
	responseMsg.SetStatusCode(404)
	responseMsg.SetBody(errorMsg)

	p.responseProcessingPipeline.Execute(responseMsg)
}
