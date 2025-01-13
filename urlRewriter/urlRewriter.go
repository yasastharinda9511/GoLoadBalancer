package urlRewriter

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
	if url, exists := rewriter.rewriteURL[ruleID]; exists {
		return url
	}
	return ""
}

func (rewriter *URLRewriter) InsertRewriteURL(ruleID string, rewriteURLPath string) {
	rewriter.rewriteURL[ruleID] = rewriteURLPath
}
