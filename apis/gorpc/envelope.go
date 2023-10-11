package gorpc

import services "resk/services"

type EnvelopeRPC struct {
}

func (e *EnvelopeRPC) SendOut(in services.RedEnvelopeSendingDTO, out *services.RedEnvelopeActivity) error {
	se := services.GetRedEnvelopeService()
	out, err := se.SendOut(in)
	if err != nil {
		return err
	}
	return nil
}
func (e *EnvelopeRPC) Receive(in services.RedEnvelopeReceiveDTO, out *services.RedEnvelopeItemDTO) error {
	se := services.GetRedEnvelopeService()
	out, err := se.Receive(in)
	if err != nil {
		return err
	}
	return nil
}
