package volumetransfers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
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

type commonResult struct {
	golangsdk.Result
}

// Extract will get the Transfer object out of the commonResult object.
func (r commonResult) Extract() (*Transfer, error) {
	var s Transfer
	err := r.ExtractInto(&s)
	return &s, err
}

// ExtractInto converts our response data into a transfer struct
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "transfer")
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	golangsdk.ErrResult
}

// ExtractTransfers extracts and returns Transfers. It is used while iterating over a transfers.List call.
func ExtractTransfers(r pagination.Page) ([]Transfer, error) {
	var s []Transfer
	err := ExtractTransfersInto(r, &s)
	return s, err
}

// ExtractTransfersInto similar to ExtractInto but operates on a `list` of transfers
func ExtractTransfersInto(r pagination.Page, v interface{}) error {
	return r.(TransferPage).Result.ExtractIntoSlicePtr(v, "transfers")
}

// TransferPage is a pagination.pager that is returned from a call to the List function.
type TransferPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no Transfers.
func (r TransferPage) IsEmpty() (bool, error) {
	transfers, err := ExtractTransfers(r)
	return len(transfers) == 0, err
}

func (page TransferPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"transfers_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}
