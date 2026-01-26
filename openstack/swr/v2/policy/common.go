package policy

type Rule struct {
	// Retention policy type.
	// The value can be "date_rule" (number of days) or "tag_rule" (number of tags).
	Template string `json:"template" required:"true"`
	// Parameters for the retention rule.
	// If Template is "date_rule", Params should contain {"days": "xxx"}.
	// If Template is "tag_rule", Params should contain {"num": "xxx"}.
	Params map[string]string `json:"params" required:"true"`
	// Exception image selectors that should not be affected by retention.
	TagSelectors []TagSelector `json:"tag_selectors" required:"true"`
}

type TagSelector struct {
	// Matching rule type for tags.
	// The value can be "label" or "regexp".
	Kind string `json:"kind" required:"true"`
	// If Kind is "label", the tag to match.
	// If Kind is "regexp", the regular expression to match.
	Pattern string `json:"pattern" required:"true"`
}
