package notifier

import "fmt"

type Notifier interface {
	Send(userid, msg string) error
}

type UserService struct {
	Notifier Notifier
}

func (s *UserService) NotifyUser(userid, msg string) error {
	if userid == "" {
		return fmt.Errorf("userid ist leer")
	}
	return s.Notifier.Send(userid, msg)
}
