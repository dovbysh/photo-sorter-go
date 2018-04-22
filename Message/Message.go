package Message

import "fmt"

// Logger should log message
type Logger interface {
	Log(msg string)
}

// LoggerImpl is implementation to Logger that print message to stdout
type LoggerImpl struct{}

// Log msg to stdout
func (l *LoggerImpl) Log(msg string) {
	fmt.Println(msg)
}

// Message is messaging system to user
type Message struct {
	L Logger
}

// NewMessage creates default Message
func NewMessage(l Logger) *Message {
	m := &Message{L: l}
	if l == nil {
		m.L = new(LoggerImpl)
	}
	return m
}

// SuccessCopied prints message for success copied files
func (m *Message) SuccessCopied(sourceFile string, destinationFile string) {
	m.L.Log("[OK] " + sourceFile + " -> " + destinationFile)
}
