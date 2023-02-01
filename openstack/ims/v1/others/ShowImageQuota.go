package others

// GET /v1/cloudimages/quota

type ShowImageQuotaResponse struct {
	Quotas Quota `json:"quotas,omitempty"`
}

type Quota struct {
	Resources []QuotaInfo `json:"resources"`
}

type QuotaInfo struct {
	// Specifies the type of the resource to be queried.
	Type string `json:"type"`
	// Specifies the used quota.
	Used int `json:"used"`
	// Specifies the total quota.
	Quota int `json:"quota"`
	// Specifies the minimum quota.
	Min int `json:"min"`
	// Specifies the maximum quota.
	Max int `json:"max"`
}
