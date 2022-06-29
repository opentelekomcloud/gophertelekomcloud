package checkpoint

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type checkpointResult struct {
	golangsdk.Result
}

type CreateResult struct {
	checkpointResult
}

type GetResult struct {
	checkpointResult
}

type Checkpoint struct {
	CreatedAt string    `json:"created_at"`
	ID        string    `json:"id"`
	ProjectId string    `json:"project_id"`
	Status    string    `json:"status"`
	Vault     Vault     `json:"vault"`
	ExtraInfo ExtraInfo `json:"extra_info"`
}

type Vault struct {
	ID               string                `json:"id"`
	Name             string                `json:"name"`
	Resources        []CheckpointResources `json:"resources"`
	SkippedResources []SkippedResources    `json:"skipped_resources"`
}

type ExtraInfo struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	RetentionDuration int    `json:"retention_duration"`
}

type CheckpointResources struct {
	ExtraInfo     string `json:"extra_info"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	ProtectStatus string `json:"protect_status"`
	ResourceSize  string `json:"resource_size"`
	Type          string `json:"type"`
	BackupSize    string `json:"backup_size"`
	BackupCount   string `json:"backup_count"`
}

type SkippedResources struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Reason string `json:"reason"`
}

func (r checkpointResult) Extract() (*Checkpoint, error) {
	s := new(Checkpoint)
	err := r.ExtractIntoStructPtr(s, "checkpoint")
	if err != nil {
		return nil, err
	}
	return s, nil
}
