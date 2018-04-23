package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type loggerMock struct {
	mock.Mock
}

// Log msg to stdout
func (m *loggerMock) Log(msg string) {
	m.Called(msg)
}

func TestSuccessCopied(t *testing.T) {
	logger := new(loggerMock)
	m := NewMessage(logger)

	logger.On("Log", "[OK] /tmp/src -> /tmp/dst").Return()

	m.SuccessCopied("/tmp/src", "/tmp/dst")

	logger.AssertExpectations(t)
}

func TestNewMessageCreateDefaultLoggerImpl(t *testing.T) {
	m := NewMessage(nil)

	assert.IsType(t, new(LoggerImpl), m.L)
}
