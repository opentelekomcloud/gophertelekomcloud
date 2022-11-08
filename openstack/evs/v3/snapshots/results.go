package snapshots

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Snapshot contains all the information associated with a Cinder Snapshot.
type Snapshot struct {
	// Unique identifier.
	ID string `json:"id"`
	// Date created.
	CreatedAt time.Time `json:"-"`
	// Date updated.
	UpdatedAt time.Time `json:"-"`
	// Display name.
	Name string `json:"name"`
	// Display description.
	Description string `json:"description"`
	// ID of the Volume from which this Snapshot was created.
	VolumeID string `json:"volume_id"`
	// Currect status of the Snapshot.
	Status string `json:"status"`
	// Size of the Snapshot, in GB.
	Size int `json:"size"`
	// User-defined key-value pairs.
	Metadata map[string]string `json:"metadata"`
}

// UnmarshalJSON converts our JSON API response into our snapshot struct
func (r *Snapshot) UnmarshalJSON(b []byte) error {
	type tmp Snapshot
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
		UpdatedAt golangsdk.JSONRFC3339MilliNoZ `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Snapshot(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return err
}

func extra(err error, raw *http.Response) (*Snapshot, error) {
	if err != nil {
		return nil, err
	}

	var res Snapshot
	err = extract.IntoStructPtr(raw.Body, &res, "snapshot")
	return &res, err
}

// WaitForStatus will continually poll the resource, checking for a particular
// status. It will do this for the amount of seconds defined.
func WaitForStatus(c *golangsdk.ServiceClient, id, status string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := Get(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == status {
			return true, nil
		}

		return false, nil
	})
}
