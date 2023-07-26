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
	ID           string
	CheckpointID string `q:"checkpoint_id"`
	ImageType    string `q:"image_type"`
	Limit        string `q:"limit"`
	Marker       string `q:"marker"`
	Name         string `q:"name"`
	Offset       string `q:"offset"`
	ParentID     string `q:"parent_id"`
	ResourceAZ   string `q:"resource_az"`
	ResourceID   string `q:"resource_id"`
	ResourceName string `q:"resource_name"`
	ResourceType string `q:"resource_type"`
	Sort         string `q:"sort"`
	Status       string `q:"status"`
	VaultID      string `q:"vault_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Backup, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("backups") + q.String(),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return BackupPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}.AllPages()
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
	m := map[string]any{}
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
