package volumetransfers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Transfer represents a Volume Transfer record
type Transfer struct {
	// Specifies the disk transfer ID.
	ID string `json:"id"`
	// Specifies the authentication key of the disk transfer.
	AuthKey string `json:"auth_key"`
	// Specifies the disk transfer name.
	Name string `json:"name"`
	// Specifies the disk ID.
	VolumeID string `json:"volume_id"`
	// Specifies the time when the disk transfer was created.
	// Time format: UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	CreatedAt time.Time `json:"-"`
	// Specifies the links of the disk transfer.
	Links []map[string]string `json:"links"`
}

// UnmarshalJSON is our unmarshalling helper
func (r *Transfer) UnmarshalJSON(b []byte) error {
	type tmp Transfer
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Transfer(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)

	return err
}

func extra(err error, raw *http.Response) (*Transfer, error) {
	if err != nil {
		return nil, err
	}

	var res Transfer
	err = extract.IntoStructPtr(raw.Body, &res, "transfer")
	return &res, err
}
