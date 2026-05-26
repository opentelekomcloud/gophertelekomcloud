package snapshots

import (
	"encoding/json"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type Snapshot struct {
	ID          string            `json:"id"`
	CreatedAt   time.Time         `json:"-"`
	UpdatedAt   time.Time         `json:"-"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	VolumeID    string            `json:"volume_id"`
	Status      string            `json:"status"`
	Size        int               `json:"size"`
	Metadata    map[string]string `json:"metadata"`
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

	return nil
}

type snapshotResponse struct {
	Snapshot Snapshot `json:"snapshot"`
}
