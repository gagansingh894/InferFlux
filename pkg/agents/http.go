package agents

type HTTPAgent struct{}

func NewHTTPAgent() (*HTTPAgent, error) {
	return &HTTPAgent{}, nil
}
