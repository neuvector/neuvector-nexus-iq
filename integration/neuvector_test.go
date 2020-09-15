package integration

import "github.com/neuvector/neuvector-nexus-iq/neuvector"

type NeuVectorNexusIqMock struct {
	ProcessWebhookRequestFunc func(*neuvector.WebhookRequest) error
}

var _ neuVectorNexusIq = &NeuVectorNexusIqMock{}

func (n *NeuVectorNexusIqMock) ProcessWebhookRequest(wr *neuvector.WebhookRequest) error {
	return n.ProcessWebhookRequestFunc(wr)
}
