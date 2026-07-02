package projects

// EnterpriseProject represents an EPS enterprise project.
type EnterpriseProject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Status      int    `json:"status"`
	DeleteFlag  bool   `json:"delete_flag"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
