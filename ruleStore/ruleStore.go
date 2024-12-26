package ruleStore

// RuleStore holds a map of rule lists indexed by an ID

import (
	"github.com/yasastharinda9511/go_gateway_api/message"
	"github.com/yasastharinda9511/go_gateway_api/rules"
)

type RuleStore struct {
	rules map[string][]rules.Rule
}

// NewRuleStore creates a new RuleStore
func NewRuleStore() *RuleStore {
	return &RuleStore{
		rules: make(map[string][]rules.Rule),
	}
}

// AddRule adds a rule to the store under the given ID
func (rs *RuleStore) AddRule(id string, rule rules.Rule) {
	rs.rules[id] = append(rs.rules[id], rule)
}

// GetRules retrieves the list of rules for the given ID
func (rs *RuleStore) GetRules(id string) []rules.Rule {
	return rs.rules[id]
}

// give me a evaloate function
func (rs *RuleStore) Evaluate(request *message.HttpRequestMessage) string {
	println("Evaluating rules")
	for id, ruleList := range rs.rules {
		eval := true
		for _, rule := range ruleList {
			if !rule.Evaluate(request) {
				eval = false
				break
			}
		}
		if eval {
			return id
		}
	}
	return ""
}

func (rs *RuleStore) PrintAllRules() {
	for id, ruleList := range rs.rules {
		println("ID:", id)
		for _, rule := range ruleList {
			rule.Print()
		}
	}
}
