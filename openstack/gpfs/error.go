package gpfs

import (
	"encoding/xml"
	"fmt"
)

// Error represents an error returned by the GPFS service.
type Error struct {
	BaseModel
	Status   string
	XMLName  xml.Name `xml:"Error"`
	Code     string   `xml:"Code"`
	Message  string   `xml:"Message"`
	Resource string   `xml:"Resource"`
	HostId   string   `xml:"HostId"`
}

func (err Error) Error() string {
	return fmt.Sprintf("gpfs: service returned error: Status=%s, Code=%s, Message=%s, RequestId=%s",
		err.Status, err.Code, err.Message, err.RequestId)
}
