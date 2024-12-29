package ruleStore

// RuleStore holds a map of rule lists indexed by an ID

import (
	"fmt"

	"github.com/yasastharinda9511/go_gateway_api/errors"
	"github.com/yasastharinda9511/go_gateway_api/message"
	"github.com/yasastharinda9511/go_gateway_api/pathtrie"
	"github.com/yasastharinda9511/go_gateway_api/rules"
)

type RuleStore struct {
	rules    map[string][]rules.Rule
	pathtrie *pathtrie.PathTrie
}

// NewRuleStore creates a new RuleStore
func NewRuleStore() *RuleStore {
	return &RuleStore{
		rules:    make(map[string][]rules.Rule),
		pathtrie: pathtrie.NewPathTrie(),
	}
}

// AddRule adds a rule to the store under the given ID
func (rs *RuleStore) AddRule(id string, rule rules.Rule) {

	if pathRule, ok := (rule).(*rules.PathRule); ok {
		rs.pathtrie.Insert(pathRule.GetPath(), id)
	}

	rs.rules[id] = append(rs.rules[id], rule)
}

// GetRules retrieves the list of rules for the given ID
func (rs *RuleStore) GetRules(id string) []rules.Rule {
	return rs.rules[id]
}

func (rs *RuleStore) Evaluate(request *message.HttpRequestMessage) (string, error) {
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
			return id, nil
		}
	}
	return "", errors.NewRuleNotFoundError(request.GetUID())
}

func (rs *RuleStore) PrintAllRules() {
	for id, ruleList := range rs.rules {
		println("ID:", id)
		for _, rule := range ruleList {
			fmt.Print(rule)
		}
	}
}
