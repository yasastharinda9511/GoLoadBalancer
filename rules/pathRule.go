package rules

import (
	"fmt"
	"strings"

	"github.com/yasastharinda9511/go_gateway_api/message"
)

type PathRule struct {
	path          string
	pathRuleParts []string
}

func NewPathRule(path string) *PathRule {
	return &PathRule{
		path:          path,
		pathRuleParts: strings.Split(path, "/"),
	}
}

// Evaluate checks if the request contains the specified header with the specified value
func (r *PathRule) Evaluate(request *message.HttpRequestMessage) bool {

	parts := strings.Split(request.GetURL(), "/")
	for i, part := range r.pathRuleParts {
		if part == "*" {
			return true
		}
		if i >= len(parts) || parts[i] != part {
			return false
		}
	}
	return len(parts) >= len(r.pathRuleParts)
}

func (r *PathRule) Print() {
	fmt.Printf("	RuleType: PathRule\n	Path: %s\n", r.path)
}

func (r *PathRule) GetPath() string {
	return r.path
}
