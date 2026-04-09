package notifier

import (
	"fmt"
	"testing"
)

type MockNotifier struct {
	NotifiedUserid  string
	NotifiedMessage string
	ShouldFail      bool
}

func (m *MockNotifier) Send(userid, msg string) error {
	if m.ShouldFail {
		return fmt.Errorf("mock konnte nicht senden")
	}
	m.NotifiedUserid = userid
	m.NotifiedMessage = msg
	return nil
}

func TestUserService(t *testing.T) {

	mock := &MockNotifier{}

	service := &UserService{Notifier: mock}

	err := service.NotifyUser("user-", "Hello")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if mock.NotifiedUserid != "user-" {
		t.Errorf(
			"Expected message to be ’Hello’. but got ’%s’",
			mock.NotifiedMessage,
		)
	}
}
