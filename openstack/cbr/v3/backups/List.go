package backups

import (
	"reflect"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the flavor attributes you want to see returned.
type ListOpts struct {
	ID             string
	CheckpointID   string `q:"checkpoint_id"`
	DedicatedCloud bool   `q:"dec"`
	EndTime        string `q:"end_time"`
	ImageType      string `q:"image_type"`
	Limit          string `q:"limit"`
	Marker         string `q:"marker"`
	MemberStatus   string `q:"member_status"`
	Name           string `q:"name"`
	Offset         string `q:"offset"`
	OwningType     string `q:"own_type"`
	ParentID       string `q:"parent_id"`
	ResourceAZ     string `q:"resource_az"`
	ResourceID     string `q:"resource_id"`
	ResourceName   string `q:"resource_name"`
	ResourceType   string `q:"resource_type"`
	Sort           string `q:"sort"`
	StartTime      string `q:"start_time"`
	Status         string `q:"status"`
	UsedPercent    string `q:"used_percent"`
	VaultID        string `q:"vault_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Backup, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("backups").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.NewPager(client, client.ServiceURL(url.String()),
		func(r pagination.PageResult) pagination.Page {
			return BackupPage{pagination.LinkedPageBase{PageResult: r}}
		}).AllPages()
	if err != nil {
		return nil, err
	}

	allBackups, err := ExtractBackups(pages)
	if err != nil {
		return nil, err
	}

	return filterBackups(allBackups, opts)
}

func filterBackups(backups []Backup, opts ListOpts) ([]Backup, error) {
	var refinedBackups []Backup
	var matched bool
	m := map[string]interface{}{}
	if opts.ID != "" {
		m["ID"] = opts.ID
	}

	if len(m) > 0 && len(backups) > 0 {
		for _, backup := range backups {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&backup, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedBackups = append(refinedBackups, backup)
			}
		}

	} else {
		refinedBackups = backups
	}

	return refinedBackups, nil
}

func getStructField(v *Backup, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
