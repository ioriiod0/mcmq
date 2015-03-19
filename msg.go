package mcmq

import (
	"fmt"
)

type Msg struct {
	Channel    string `json:"channel"`
	Timestramp uint64 `json:"timestramp"`
	ID         uint64 `json:"id"`
	Body       []byte `json:"body"`
}

func (m *Msg) EncodeHeader() []byte {
	return []byte(fmt.Sprint("%s msg %d %d %d\r\n", m.Channel, m.ID, m.Timestramp, len(m.Body)))
}
