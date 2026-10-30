package certificates

type Certificate struct {
	ID          string `json:"id"`
	ProjectID   string `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Domain      string `json:"domain"`
	PrivateKey  string `json:"private_key"`
	Certificate string `json:"certificate"`
	// AdminStateUp is documented but marked unsupported by the reviewed OTC API.
	AdminStateUp bool   `json:"admin_state_up"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	ExpireTime   string `json:"expire_time"`
}
