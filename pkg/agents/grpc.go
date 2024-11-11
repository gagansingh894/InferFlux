package agents

type GRPCAgent struct{}

func NewGRPCAgent() (*GRPCAgent, error) {
	return &GRPCAgent{}, nil
}
