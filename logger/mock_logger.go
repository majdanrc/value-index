package logger

type MockLogger struct{}

func NewMockLogger() *MockLogger {
	return &MockLogger{}
}

func (m *MockLogger) Debug(msg string) {}
func (m *MockLogger) Info(msg string)  {}
func (m *MockLogger) Warn(msg string)  {}
func (m *MockLogger) Error(msg string) {}
