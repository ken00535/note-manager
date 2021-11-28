package logger

import "fmt"

type mockLogger struct{}

func (logger *mockLogger) Debug(args ...interface{}) {}
func (logger *mockLogger) Info(args ...interface{}) {
	fmt.Println(args...)
}
func (logger *mockLogger) Warn(args ...interface{})  {}
func (logger *mockLogger) Error(args ...interface{}) {}
func (logger *mockLogger) Fatal(args ...interface{}) {}
func (logger *mockLogger) Panic(args ...interface{}) {}

// NewMockLogger new a mock of logger
func NewMockLogger() Logger {
	return &mockLogger{}
}
