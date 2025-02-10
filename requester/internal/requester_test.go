package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockLogger struct {
	infoMessages  []string
	errorMessages []string
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		infoMessages:  make([]string, 0),
		errorMessages: make([]string, 0),
	}
}

func (l *MockLogger) Info(msg string) {
	l.infoMessages = append(l.infoMessages, msg)
}

func (l *MockLogger) Error(msg string) {
	l.errorMessages = append(l.errorMessages, msg)
}

func (l *MockLogger) Debug(msg string) {
}

func TestRequesterSuccess(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	logger := NewMockLogger()
	params := NewParams(3, server.URL)
	requester := NewRequester(params, logger)

	// Act
	requester.Run()

	// Assert
	assert.Equal(t, 7, len(logger.infoMessages)) // 1 inicio + 3 requests + 3 responses
	assert.Equal(t, 0, len(logger.errorMessages))
	assert.Contains(t, logger.infoMessages[0], "Sending 3 requests")
}

func TestRequesterWithServerError(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	logger := NewMockLogger()
	params := NewParams(2, server.URL)
	requester := NewRequester(params, logger)

	// Act
	requester.Run()

	// Assert
	assert.Equal(t, 5, len(logger.infoMessages)) // 1 inicio + 2 requests + 2 responses
	assert.Equal(t, 0, len(logger.errorMessages))
	assert.Contains(t, logger.infoMessages[0], "Sending 2 requests")
}

func TestRequesterWithInvalidHost(t *testing.T) {
	// Arrange
	logger := NewMockLogger()
	params := NewParams(1, "http://invalid-host")
	requester := NewRequester(params, logger)

	// Act
	requester.Run()

	// Assert
	assert.Equal(t, 2, len(logger.infoMessages)) // 1 inicio + 1 request
	assert.Equal(t, 1, len(logger.errorMessages))
	assert.Contains(t, logger.infoMessages[0], "Sending 1 requests")
	assert.Contains(t, logger.errorMessages[0], "Error sending request")
}
