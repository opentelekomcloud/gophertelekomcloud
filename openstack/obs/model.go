package obs

import (
	"encoding/xml"
)

type getBucketAclOutputObs struct {
	BaseModel
	accessControlPolicyObs
}

type Error struct {
	XMLName   xml.Name `xml:"Error"`
	Key       string   `xml:"Key"`
	VersionId string   `xml:"VersionId"`
	Code      string   `xml:"Code"`
	Message   string   `xml:"Message"`
}
