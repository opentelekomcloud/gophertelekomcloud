package postgres_extensions

type PostgresExtensionOpts struct {
	// Specifies the name of the specific database created inside the RDS instance.
	// This is the logical database name, not the RDS instance identifier.
	DatabaseName string `json:"database_name"  required:"true"`
	// Specifies the extension name
	ExtensionName string `json:"extension_name"  required:"true"`
}
