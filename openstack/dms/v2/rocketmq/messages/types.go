package messages

// Property is a message attribute.
type Property struct {
	// Attribute name.
	Name string `json:"name,omitempty"`
	// Attribute value.
	Value string `json:"value,omitempty"`
}

// ResendResponse is returned by ResendDeadLetter and VerifyConsumption.
type ResendResponse struct {
	// Result for each message.
	ResendResults []ResendResult `json:"resend_results"`
}

type ResendResult struct {
	// Message ID.
	MsgID string `json:"msg_id"`
	// Error code.
	ErrorCode string `json:"error_code"`
	// Error message.
	ErrorMessage string `json:"error_message"`
}
