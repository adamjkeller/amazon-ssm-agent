package channelmock

import "github.com/stretchr/testify/mock"

type MockedChannel struct {
	mock.Mock
}

func (m *MockedChannel) Open(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockedChannel) Send(msg string) error {
	args := m.Called(msg)
	return args.Error(0)
}

func (m *MockedChannel) GetMessageChannel() chan string {
	args := m.Called()
	return args.Get(0).(chan string)
}

func (m *MockedChannel) Close() {
	m.Called()
	return
}
