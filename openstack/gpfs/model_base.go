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
	BucketType   string    `xml:"BucketType,omitempty"`
}

// Owner defines owner properties
type Owner struct {
	XMLName     xml.Name `xml:"Owner"`
	ID          string   `xml:"ID"`
	DisplayName string   `xml:"DisplayName,omitempty"`
}

// BucketLocation defines bucket location configuration
type BucketLocation struct {
	XMLName  xml.Name `xml:"CreateBucketConfiguration"`
	Location string   `xml:"LocationConstraint,omitempty"`
}
