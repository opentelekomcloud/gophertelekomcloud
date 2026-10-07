package gpfs

// ISseHeader defines the sse encryption header
type ISseHeader interface {
	GetEncryption() string
	GetKey() string
}

// SseCHeader defines the SseC header
type SseCHeader struct {
	Encryption string
	Key        string
	KeyMD5     string
}
