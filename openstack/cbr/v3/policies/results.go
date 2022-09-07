package policies

type Policy struct {
	ID                  string                 `json:"id"`
	Name                string                 `json:"name"`
	Enabled             bool                   `json:"enabled"`
	OperationDefinition *PolicyODCreate        `json:"operation_definition"`
	OperationType       OperationType          `json:"operation_type"`
	Trigger             *PolicyTriggerResp     `json:"trigger"`
	AssociatedVaults    []PolicyAssociateVault `json:"associated_vaults"`
}

type PolicyAssociateVault struct {
	VaultID            string `json:"vault_id"`
	DestinationVaultID string `json:"destination_vault_id"`
}

type PolicyTriggerPropertiesResp struct {
	Pattern   []string `json:"pattern"`
	StartTime string   `json:"start_time"`
}

type PolicyTriggerResp struct {
	Properties PolicyTriggerPropertiesResp `json:"properties"`
}
