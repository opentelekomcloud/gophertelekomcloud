package topics

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type CreateResult struct {
	golangsdk.Result
}

type GetResult struct {
	golangsdk.Result
}

type DeleteResult struct {
	golangsdk.Result
}

type TopicName struct {
	Name string `json:"id"`
}

type Topic struct {
	Size             int          `json:"size"`
	RemainPartitions int          `json:"remain_partitions"`
	MaxPartitions    int          `json:"max_partitions"`
	Topics           []Parameters `json:"topics"`
}

type Parameters struct {
	Name             string `json:"id"`
	Partition        int    `json:"partition"`
	Replication      int    `json:"replication"`
	SyncReplication  string `json:"sync_replication"`
	RetentionTime    int    `json:"retention_time"`
	SyncMessageFlush string `json:"sync_message_flush"`
}

func (r CreateResult) Extract() (*TopicName, error) {
	s := new(TopicName)
	err := r.ExtractIntoStructPtr(s, "")
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r GetResult) Extract() (*Topic, error) {
	s := new(Topic)
	err := r.ExtractIntoStructPtr(s, "")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type TopicDelete struct {
	Name    string `json:"id"`
	Success bool   `json:"success"`
}

func (r DeleteResult) Extract() ([]TopicDelete, error) {
	var s []TopicDelete
	err := r.ExtractIntoSlicePtr(&s, "topics")
	if err != nil {
		return nil, err
	}
	return s, nil
}
