package emailutils

// Email - Struct to send text email
type TextEmail struct {
	From     string   `json:"from,omitempty"`
	Password string   `json:"password,omitempty"`
	To       []string `json:"to,omitempty"`
	Co       []string `json:"co,omitempty"`
	Cco      []string `json:"cco,omitempty"`
	Subject  string   `json:"subject,omitempty"`
	Body     string   `json:"body,omitempty"`
}
