package ies

import "github.com/lvdund/ngap/aper"

const (
	CausePresentNothing uint64 = iota /* No components present */
	CausePresentRadioNetwork
	CausePresentTransport
	CausePresentNas
	CausePresentProtocol
	CausePresentMisc
	CausePresentChoiceExtensions
)

type Cause struct {
	Choice       uint64
	RadioNetwork *CauseRadioNetwork `False,,,`
	Transport    *CauseTransport    `False,,,`
	Nas          *CauseNas          `False,,,`
	Protocol     *CauseProtocol     `False,,,`
	Misc         *CauseMisc         `False,,,`
	// ChoiceExtensions *CauseExtIEs `False,,,`
}

func (ie *Cause) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return
	}
	switch ie.Choice {
	case CausePresentRadioNetwork:
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
	case CausePresentRadioNetwork:
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
