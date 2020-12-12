package jsession

type session struct {
}

type sessionConfig struct {
}

func NewJSession() *session {
	return &session{}
}

// Start starts the session
func (s *session) Start() {

}

// Set sets the session
func (s *session) Set(key string, value string) {

}

// Get gets the session
func (s *session) Get(key string) string {
	return ""
}

// Expire expires the session
func (s *session) Expire() {

}
