package structs

type ChargeInfo struct {
	ChargeMode  string `json:"charge_mode" required:"true"`
	PeriodType  string `json:"period_type,omitempty"`
	PeriodNum   int    `json:"period_num,omitempty"`
	IsAutoRenew string `json:"is_auto_renew,omitempty"`
	IsAutoPay   string `json:"is_auto_pay,omitempty"`
}

type ResourceRef struct {
	ID string `json:"id"`
}

func ExtractIDs(refs []ResourceRef) []string {
	ids := make([]string, len(refs))
	for i, v := range refs {
		ids[i] = v.ID
	}
	return ids
}
