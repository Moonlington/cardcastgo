package cardcastgo

func New(token string) (s *Session, err error) {
	s = &Session{}
	s.Token = token
	return
}
