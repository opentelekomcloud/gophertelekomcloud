package rules

type RuleType string
type CompareType string

const (
	HostName RuleType = "HOST_NAME"
	Path     RuleType = "PATH"

	EqualTo    CompareType = "EQUAL_TO"
	Regex      CompareType = "REGEX"
	StartsWith CompareType = "STARTS_WITH"
)

type Condition struct {
	// Specifies the key of match item.
	Key string `json:"key"`

	// Specifies the value of the match item.
	Value string `json:"value" required:"true"`
}
