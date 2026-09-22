package alias

type VectorStrategy struct {
	CombineType string                `json:"combine_type"`
	Rules       *VersionStrategyRules `json:"rules"`
}

type VersionStrategyRules struct {
	RuleType string `json:"rule_type"`
	Param    string `json:"param"`
	Op       string `json:"op"`
	Value    string `json:"value"`
}

// ####### RESPONSE STRUCTURES ########

type FuncAliasesResp struct {
	Name                      string                    `json:"name"`
	Version                   string                    `json:"version"`
	Description               string                    `json:"description"`
	LastModified              string                    `json:"last_modified"`
	AliasUrn                  string                    `json:"alias_urn"`
	AdditionalVersionWeights  map[string]int            `json:"additional_version_weights"`
	AdditionalVersionStrategy map[string]VectorStrategy `json:"additional_version_strategy"`
}
