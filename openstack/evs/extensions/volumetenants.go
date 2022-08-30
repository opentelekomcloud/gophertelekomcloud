package extensions

type VolumeTenantExt struct {
	// TenantID is the id of the project that owns the volume.
	TenantID string `json:"os-vol-tenant-attr:tenant_id"`
}
