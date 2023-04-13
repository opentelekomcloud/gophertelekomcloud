package monitors

type Type string

// Constants that represent approved monitoring types.
const (
	TypePING  Type = "PING"
	TypeTCP   Type = "TCP"
	TypeHTTP  Type = "HTTP"
	TypeHTTPS Type = "HTTPS"
)
