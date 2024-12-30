package ies

import "github.com/lvdund/ngap/aper"

const (
	CausePresentNothing uint64 = iota
	CausePresentRadionetwork
	CausePresentTransport
	CausePresentNas
	CausePresentProtocol
	CausePresentMisc
	CausePresentChoiceExtensions
)

type Cause struct {
	Choice       uint64
	RadioNetwork *CauseRadioNetwork
	Transport    *CauseTransport
	Nas          *CauseNas
	Protocol     *CauseProtocol
	Misc         *CauseMisc
	// ChoiceExtensions *CauseExtIEs
}

func (ie *Cause) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return
	}
	switch ie.Choice {
	case CausePresentRadionetwork:
		err = ie.RadioNetwork.Encode(w)
	case CausePresentTransport:
		err = ie.Transport.Encode(w)
	case CausePresentNas:
		err = ie.Nas.Encode(w)
	case CausePresentProtocol:
		err = ie.Protocol.Encode(w)
	case CausePresentMisc:
		err = ie.Misc.Encode(w)
	}
	return
}
func (ie *Cause) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return
	}
	switch ie.Choice {
	case CausePresentRadionetwork:
		var tmp CauseRadioNetwork
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RadioNetwork = &tmp
	case CausePresentTransport:
		var tmp CauseTransport
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Transport = &tmp
	case CausePresentNas:
		var tmp CauseNas
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Nas = &tmp
	case CausePresentProtocol:
		var tmp CauseProtocol
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Protocol = &tmp
	case CausePresentMisc:
		var tmp CauseMisc
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Misc = &tmp
	}
	return
}
