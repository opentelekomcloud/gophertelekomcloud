package gpfs

import (
	"encoding/xml"
	"time"
)

// Bucket defines bucket properties
type Bucket struct {
	XMLName      xml.Name  `xml:"Bucket"`
	Name         string    `xml:"Name"`
	CreationDate time.Time `xml:"CreationDate"`
	Location     string    `xml:"Location"`
}

// Owner defines owner properties
type Owner struct {
	ID string `xml:"ID"`
}

// BucketLocation defines bucket location configuration
type BucketLocation struct {
	XMLName  xml.Name `xml:"CreateBucketConfiguration"`
	Location string   `xml:"Location,omitempty"`
}
