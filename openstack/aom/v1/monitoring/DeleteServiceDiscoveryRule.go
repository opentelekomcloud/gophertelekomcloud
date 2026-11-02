package monitoring

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DeleteServiceDiscoveryRuleOpts struct {
	// Discovery rule IDs, separated by commas (,). The parameter cannot be empty.
	AppRulesIds string `q:"appRulesIds,required"`
}

// DeleteServiceDiscoveryRule is used to delete a service discovery rule.
// Multiple rule IDs are sent as a comma-separated query parameter and cannot be empty.
func DeleteServiceDiscoveryRule(client *golangsdk.ServiceClient, ruleIds []string) (*ServiceDiscoveryRuleOperationResp, error) {
	opts := DeleteServiceDiscoveryRuleOpts{
		AppRulesIds: strings.Join(ruleIds, ","),
	}

	// DELETE /v1/{project_id}/inv/servicediscoveryrules?appRulesIds=<id1>,<id2>
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("inv", "servicediscoveryrules").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// DELETE /v1/{project_id}/inv/servicediscoveryrules?appRulesIds=<id1>,<id2>
	// The API doc marks Content-Type as mandatory; since there is no request
	// body, golangsdk does not infer the header, so set it explicitly.
	raw, err := client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ServiceDiscoveryRuleOperationResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
