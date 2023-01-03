package obs

// ISseHeader defines the sse encryption header
type ISseHeader interface {
	GetEncryption() string
	GetKey() string
}

// SseKmsHeader defines the SseKms header
type SseKmsHeader struct {
	Encryption string
	Key        string
	isObs      bool
}

// SseCHeader defines the SseC header
type SseCHeader struct {
	Encryption string
	Key        string
	KeyMD5     string
}
