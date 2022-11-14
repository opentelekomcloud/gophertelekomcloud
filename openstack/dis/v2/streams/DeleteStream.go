package streams

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteStream(client *golangsdk.ServiceClient, streamName string) (err error) {
	// DELETE /v2/{project_id}/streams/{stream_name}
	_, err = client.Delete(client.ServiceURL("streams", streamName), nil)
	return
}
