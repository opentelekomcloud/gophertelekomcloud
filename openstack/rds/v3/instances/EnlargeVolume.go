package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type EnlargeVolumeRdsOpts struct {
	//
	EnlargeVolume *EnlargeVolumeSize `json:"enlarge_volume" required:"true"`
}

type EnlargeVolumeSize struct {
	//
	Size int `json:"size" required:"true"`
}

type EnlargeVolumeBuilder interface {
	ToEnlargeVolumeMap() (map[string]interface{}, error)
}

func (opts EnlargeVolumeRdsOpts) ToEnlargeVolumeMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func EnlargeVolume(client *golangsdk.ServiceClient, opts EnlargeVolumeBuilder, instanceId string) (r EnlargeVolumeResult) {
	b, err := opts.ToEnlargeVolumeMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("instances", instanceId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})

	return
}
