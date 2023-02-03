package aggregates

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Aggregate represents a host aggregate in the OpenStack cloud.
type Aggregate struct {
	// The availability zone of the host aggregate.
	AvailabilityZone string `json:"availability_zone"`
	// A list of host ids in this aggregate.
	Hosts []string `json:"hosts"`
	// The ID of the host aggregate.
	ID int `json:"id"`
	// Metadata key and value pairs associate with the aggregate.
	Metadata map[string]string `json:"metadata"`
	// Name of the aggregate.
	Name string `json:"name"`
	// The date and time when the resource was created.
	CreatedAt time.Time `json:"-"`
	// The date and time when the resource was updated,
	// if the resource has not been updated, this field will show as null.
	UpdatedAt time.Time `json:"-"`
	// The date and time when the resource was deleted,
	// if the resource has not been deleted yet, this field will be null.
	DeletedAt time.Time `json:"-"`
	// A boolean indicates whether this aggregate is deleted or not,
	// if it has not been deleted, false will appear.
	Deleted bool `json:"deleted"`
}

// UnmarshalJSON to override default
func (r *Aggregate) UnmarshalJSON(b []byte) error {
	type tmp Aggregate
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
		UpdatedAt golangsdk.JSONRFC3339MilliNoZ `json:"updated_at"`
		DeletedAt golangsdk.JSONRFC3339MilliNoZ `json:"deleted_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Aggregate(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)
	r.DeletedAt = time.Time(s.DeletedAt)

	return nil
}

func extra(err error, raw *http.Response) (*Aggregate, error) {
	if err != nil {
		return nil, err
	}

	var res Aggregate
	err = extract.IntoStructPtr(raw.Body, &res, "aggregate")
	return &res, err
}
