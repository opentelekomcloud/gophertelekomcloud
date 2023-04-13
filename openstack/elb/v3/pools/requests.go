package pools

// SessionPersistence represents the session persistence feature of the load
// balancing service. It attempts to force connections or requests in the same
// session to be processed by the same member as long as it is active. Three
// types of persistence are supported:
type SessionPersistence struct {
	// The type of persistence mode.
	Type string `json:"type" required:"true"`

	// Name of cookie if persistence mode is set appropriately.
	CookieName string `json:"cookie_name,omitempty"`

	// PersistenceTimeout specifies the stickiness duration, in minutes.
	PersistenceTimeout int `json:"persistence_timeout,omitempty"`
}

type SlowStart struct {
	// Specifies whether to Enable slow start.
	Enable bool `json:"enable" required:"true"`

	// Specifies the slow start Duration, in seconds.
	Duration int `json:"duration" required:"true"`
}
