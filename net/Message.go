package net

type Message struct {
	Id     uint32
	MsgLen uint32
	Data   []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id:     id,
		MsgLen: uint32(len(data)),
		Data:   data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}
func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}
func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}
func (m *Message) SetMsgLen(msgLen uint32) {
	m.MsgLen = msgLen

}
func (m *Message) SetData(data []byte) {
	m.Data = data
}
