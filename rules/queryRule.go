package rules

import (
	"fmt"

	"github.com/yasastharinda9511/go_gateway_api/message"
)

type QueryRule struct {
	query      string
	queryValue string
}

func NewQueryRule(query string, queryValue string) *QueryRule {
	return &QueryRule{
		query:      query,
		queryValue: queryValue,
	}
}

// Evaluate checks if the request contains the specified header with the specified value
func (r *QueryRule) Evaluate(request *message.HttpRequestMessage) bool {
	return request.GetQueryParams()[r.query] == r.queryValue
}

func (r *QueryRule) Print() {
	fmt.Printf("	RuleType: QueryRule\n	Query: %s\n	QueryValue: %s\n", r.query, r.queryValue)
}
