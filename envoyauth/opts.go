package envoyauth

// An Opt allowing modifying decision ID
type Opt func(*EvalResult)

// WithDecisionID  sets the externally sent decision ID
func WithDecisionID(decisionID string) Opt {
	return func(result *EvalResult) {
		result.DecisionID = decisionID
	}
}
