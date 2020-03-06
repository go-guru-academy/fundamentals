package hub

type Hub struct {
	messages chan []byte
}

const (
	MESSAGE_LIMIT = 10
)

var (
	hub *Hub
)

func Init() {
	hub = &Hub{
		messages: make(chan []byte, MESSAGE_LIMIT),
	}
}

func Enqueue(message []byte) {
	hub.messages <- message
}

func Dequeue() []byte {
	return <-hub.messages
}
