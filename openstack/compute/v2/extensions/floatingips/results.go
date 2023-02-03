package floatingips

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// A FloatingIP is an IP that can be associated with a server.
type FloatingIP struct {
	// ID is a unique ID of the Floating IP
	ID string `json:"-"`
	// FixedIP is a specific IP on the server to pair the Floating IP with.
	FixedIP string `json:"fixed_ip,omitempty"`
	// InstanceID is the ID of the server that is using the Floating IP.
	InstanceID string `json:"instance_id"`
	// IP is the actual Floating IP.
	IP string `json:"ip"`
	// Pool is the pool of Floating IPs that this Floating IP belongs to.
	Pool string `json:"pool"`
}

func (r *FloatingIP) UnmarshalJSON(b []byte) error {
	type tmp FloatingIP
	var s struct {
		tmp
		ID interface{} `json:"id"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = FloatingIP(s.tmp)

	switch t := s.ID.(type) {
	case float64:
		r.ID = strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		r.ID = t
	}

	return err
}

func extra(err error, raw *http.Response) (*FloatingIP, error) {
	if err != nil {
		return nil, err
	}

	var res FloatingIP
	err = extract.IntoStructPtr(raw.Body, &res, "floating_ip")
	return &res, err
}
