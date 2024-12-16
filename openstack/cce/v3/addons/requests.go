package addons

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json"},
}

// Delete will permanently delete a particular addon based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id, clusterId string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id, clusterId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

type ListOptsBuilder interface {
	ToAddonListQuery() (string, error)
}

type ListOpts struct {
	Name string `q:"addon_template_name"`
}

func (opts ListOpts) ToAddonListQuery() (string, error) {
	u, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func ListTemplates(c *golangsdk.ServiceClient, clusterID string, opts ListOptsBuilder) (r ListTemplateResult) {
	url := templatesURL(c, clusterID)
	if opts != nil {
		q, err := opts.ToAddonListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += q
	}
	_, r.Err = c.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func GetTemplates(c *golangsdk.ServiceClient) (r ListTemplateResult) {
	_, r.Err = c.Get(addonTemplatesURL(c), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func ListAddonInstances(c *golangsdk.ServiceClient, clusterID string) (r ListInstanceResult) {
	_, r.Err = c.Get(instanceURL(c, clusterID), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// WaitForAddonRunning - wait until addon status is `running`
func WaitForAddonRunning(client *golangsdk.ServiceClient, id, clusterID string, timeoutSeconds int) error {
	return golangsdk.WaitFor(timeoutSeconds, func() (bool, error) {
		addon, err := Get(client, id, clusterID)
		if err != nil {
			return false, fmt.Errorf("error retriving addon status: %w", err)
		}
		if addon.Status.Status == "running" {
			return true, nil
		}
		return false, nil
	})
}

// WaitForAddonDeleted - wait until addon is deleted
func WaitForAddonDeleted(client *golangsdk.ServiceClient, id, clusterID string, timeoutSeconds int) error {
	return golangsdk.WaitFor(timeoutSeconds, func() (bool, error) {
		_, err := Get(client, id, clusterID)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, fmt.Errorf("error retriving addon status: %w", err)
		}
		return false, nil
	})
}
