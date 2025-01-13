package urlRewriter

import "fmt"

type URLRewriter struct {
	rewriteURL map[string]string
}

// NewRuleStore creates a new RuleStore
func NewURLRewriter() *URLRewriter {
	return &URLRewriter{
		rewriteURL: make(map[string]string),
	}
}

func (rewriter *URLRewriter) GetRewriteURL(ruleID string) string {
	fmt.Println("ruleID is ", ruleID)
	if url, exists := rewriter.rewriteURL[ruleID]; exists {
		fmt.Println("url is ", url)
		return url
	}
	return ""
}

func (rewriter *URLRewriter) InsertRewriteURL(ruleID string, rewriteURLPath string) {
	rewriter.rewriteURL[ruleID] = rewriteURLPath
}
