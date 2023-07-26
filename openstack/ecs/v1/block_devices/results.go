package block_devices

import "github.com/opentelekomcloud/gophertelekomcloud"

type VolumeAttachment struct {
	PciAddress string `json:"pciAddress"`
	Size       int    `json:"size"`
	BootIndex  int    `json:"bootIndex"`
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*VolumeAttachment, error) {
	s := &VolumeAttachment{}
	return s, r.ExtractInto(s)
}

func (r GetResult) ExtractInto(v any) error {
	return r.ExtractIntoStructPtr(v, "volumeAttachment")
}
