package public

type BatchReplaceDefinerReq struct {

	// 批量设置replaceDefiner请求列表
	Jobs []ReplaceDefinerInfo `json:"jobs"`
}

type ReplaceDefinerInfo struct {

	// 任务id
	JobId string `json:"job_id"`

	// 是否使用目标库的用户替换掉definer
	ReplaceDefiner bool `json:"replace_definer"`
}

// POST /v3/{project_id}/jobs/batch-replace-definer

// BatchJobsResponse
