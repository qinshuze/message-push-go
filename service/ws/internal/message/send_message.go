package message

const (
	Join  = "join"
	Leave = "leave"
	Relay = "relay"
)

type SendMessage struct {
	Type    string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
	Sender  string `json:"sender,omitempty"`
}
