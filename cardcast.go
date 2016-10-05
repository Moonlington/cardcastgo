package cardcastgo

import (
	"fmt"
)

func New(token string) (s *Session, err error) {
	s = &Session{
		Token: "",
	}

	switch t := token.(type) {
	case string:
		s.Token = t
		return
	default:
		err = fmt.Errorf("Unsupported parameter type provided")
		return
	}
	return
}
