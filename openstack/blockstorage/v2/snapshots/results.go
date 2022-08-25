package snapshots

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

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
