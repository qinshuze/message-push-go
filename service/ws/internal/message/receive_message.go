package message

import (
	"encoding/json"
)

type ReceiveMessage struct {
	roomIds []string
	names   []string
	tags    []string
	content string
}

type jsonReceiveMessage struct {
	RoomIds []string `json:"roomIds,omitempty"`
	Names   []string `json:"names,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Content string   `json:"content,omitempty"`
}

func (r *ReceiveMessage) RoomIds() []string {
	return r.roomIds
}

func (r *ReceiveMessage) Names() []string {
	return r.names
}

func (r *ReceiveMessage) Tags() []string {
	return r.tags
}

func (r *ReceiveMessage) Content() string {
	return r.content
}

func NewReceiveMessageByJsonStr(str string) (*ReceiveMessage, error) {
	var jMap jsonReceiveMessage
	err := json.Unmarshal([]byte(str), &jMap)
	if err != nil {
		return nil, err
	}

	return NewReceiveMessage(jMap.RoomIds, jMap.Names, jMap.Tags, jMap.Content), nil
}

func NewReceiveMessage(roomIds []string, names []string, tags []string, content string) *ReceiveMessage {
	receiveMsg := &ReceiveMessage{
		roomIds: roomIds,
		names:   names,
		tags:    tags,
		content: content,
	}

	return receiveMsg
}
