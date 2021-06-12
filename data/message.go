package data

const MaxLength = 2000


type MessageReader interface {
	readMessage() string
}

type MessageReset interface {
	resetMessage()
}

type MessageAdder interface {
	addMessage(string)
}


type Message struct {
	message string
}


func (m *Message)  ReadMessage() string {
	return m.message
}

func (m *Message)  ResetMessage() {
	m.message = ""
}

func (m *Message)  AddMessage(toAdd string) {
	newMessage := m.message + toAdd
	msgLen := len(newMessage)
	if msgLen > 2000 {
		newMessage = newMessage[len(newMessage)-2000:]
	}
	m.message = newMessage
}
