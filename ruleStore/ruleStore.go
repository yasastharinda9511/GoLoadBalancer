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

	// Get the rule IDs that match the request path

	ruleIDs := rs.pathtrie.MatchExactPaths(request.GetURL())

	//fmt.Printf("match exact path count %d", len(ruleIDs))
	for _, ruleID := range ruleIDs {
		if rs.evaluateRuleID(ruleID, request) {
			fmt.Printf("Exact Match %s\n", ruleID)
			return ruleID, nil
		}
	}

	ruleIDs = rs.pathtrie.MatchPrefixPaths(request.GetURL())
	//fmt.Printf("match prefix path count %d", len(ruleIDs))
	for _, ruleID := range ruleIDs {
		if rs.evaluateRuleID(ruleID, request) {
			//fmt.Printf("Exact Match %s\n", ruleID)
			return ruleID, nil
		}
	}

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

func (rs *RuleStore) evaluateRuleID(ruleID string, request *message.HttpRequestMessage) bool {
	if ruleList, exists := rs.rules[ruleID]; exists {
		for _, rule := range ruleList {
			if !rule.Evaluate(request) {
				return false
			}
		}
		return true
	}
	return false
}
