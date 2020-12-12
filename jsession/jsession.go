package jsession

type jsession struct {
}

type config struct {
}

func NewJSession() *jsession {
	return &jsession{}
}

// Start starts the session
func (s *jsession) Start() {

}

// Set sets the session
func (s *jsession) Set(key string, value string) {

}

// Get gets the session
func (s *jsession) Get(key string) string {
	return ""
}

// Expire expires the session
func (s *jsession) Expire() {

}
