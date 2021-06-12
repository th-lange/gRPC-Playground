package data

const MaxLength = 2000


type MessageReader interface {
	ReadMessage() string
}

type MessageReset interface {
	ResetMessage()
}

type MessageAdder interface {
	AddMessage(string)
}

type MessageSetter interface {
	SetMessage(string)
}

type Message struct {
	message string
}


func (m *Message)  SetMessage(message string) {
	msgLen := len(message)
	if msgLen > 2000 {
		message = message[msgLen-2000:]
	}
	m.message = message
}


func (m *Message)  ReadMessage() string {
	return m.message
}

func (m *Message)  ResetMessage() {
	m.message = ""
}

func (m *Message)  AddMessage(toAdd string) {
	m.SetMessage(m.message + toAdd)
}
