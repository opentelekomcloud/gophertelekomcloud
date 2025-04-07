package dns

type DomainSetInfoDto struct {
	// Domain name, for example, www.test.com.
	DomainName string `json:"domain_name" required:"true"`
	// Domain name description.
	Description string `json:"description,omitempty"`
}

type GetDomainNameGroupListQueryParams struct {
	// Firewall ID.
	FwInstanceID string `q:"fw_instance_id" required:"true"`
	// Offset, which specifies the start position of the record to be returned. The value must be a number no less than 0. The default value is 0.
	Offset string `q:"offset" required:"true"`
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `q:"limit" required:"true"`
	// Keyword, which can be the domain name group name or description.
	Keyword string `q:"key_word,omitempty"`
}

type GetDomainNameListQueryParams struct {
	// Firewall ID.
	FwInstanceID string `q:"fw_instance_id" required:"true"`
	// Offset, which specifies the start position of the record to be returned. The value must be a number no less than 0. The default value is 0.
	Offset string `q:"offset" required:"true"`
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `q:"limit" required:"true"`
	// Domain name, for example, www.test.com.
	DomainName string `q:"domain_name,omitempty"`
}

/*
########################## RESPONSE STRUCTS ##############################
*/

type CommonDomainNameGroupDataResponse struct {
	// Returned data for adding a domain name group.
	Data DomainSetResponseData `json:"data"`
}

type DomainSetResponseData struct {
	// Domain name group ID.
	Id string `json:"id"`
	// Domain name group name.
	Name string `json:"name"`
}

type GetDomainNameGroupDataResponse struct {
	//Returned data for querying the domain name group list.
	Data ListDomainSetsResponseData `json:"data"`
}

type ListDomainSetsResponseData struct {
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `json:"limit"`
	// Offset, which specifies the start position of the record to be returned.
	// The value must be a number no less than 0. The default value is 0.
	Offset int `json:"offset"`
	// Total number of domain names.
	Total int `json:"total"`
	// Domain name information list.
	Records []DomainSetVO `json:"records"`
}

type DomainSetVO struct {
	// Domain name group ID.
	SetID string `json:"set_id"`
	// Domain name group name.
	Name string `json:"name"`
	// Domain name group description.
	Description string `json:"description"`
	// Number of times a domain name group is referenced by rules.
	RefCount int `json:"ref_count"`
	// Domain name group type:
	// 0 - application domain name group,
	// 1 - network domain name group.
	DomainSetType int `json:"domain_set_type"`
	// Configuration status:
	// -1 - unconfigured,
	//  0 - configuration failed,
	//  1 - configuration succeeded,
	//  2 - configuring,
	//  3 - normal,
	//  4 - abnormal.
	ConfigStatus int `json:"config_status"`
	// Used rule ID list.
	Rules []UseRuleVO `json:"rules"`
}

type UseRuleVO struct {
	// Rule ID.
	ID string `json:"id"`

	// Rule name.
	Name string `json:"name"`
}

type GetDomainNamesDataResponse struct {
	// ListDomainResponseData represents the response data for listing domain names in domain name group.
	Data ListDomainResponseData `json:"data"`
}

type ListDomainResponseData struct {
	// Number of records displayed on each page. The value ranges from 1 to 1024.
	Limit int `json:"limit"`
	// Offset, which specifies the start position of the record to be returned.
	// The value must be a number no less than 0. The default value is 0.
	Offset int `json:"offset"`
	// Project ID.
	ProjectID string `json:"project_id"`
	// Domain name information list.
	Records []DomainInfo `json:"records"`
	// Domain name group ID.
	SetID string `json:"set_id"`
	// Total number of domain names.
	Total int `json:"total"`
}

type DomainInfo struct {
	// Domain name ID.
	DomainAddressID string `json:"domain_address_id"`
	// Domain name, for example, www.test.com.
	DomainName string `json:"domain_name"`
	// Domain name description.
	Description string `json:"description"`
}
