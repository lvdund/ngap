package ies

import "github.com/lvdund/ngap/aper"

const (
	CausePresentNothing int = iota /* No components present */
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
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.RadioNetwork.Encode(w)
	case 2:
		err = ie.Transport.Encode(w)
	case 3:
		err = ie.Nas.Encode(w)
	case 4:
		err = ie.Protocol.Encode(w)
	case 5:
		err = ie.Misc.Encode(w)
	}
	return
}
func (ie *Cause) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp CauseRadioNetwork
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RadioNetwork = &tmp
	case 2:
		var tmp CauseTransport
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Transport = &tmp
	case 3:
		var tmp CauseNas
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Nas = &tmp
	case 4:
		var tmp CauseProtocol
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Protocol = &tmp
	case 5:
		var tmp CauseMisc
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Misc = &tmp
	}
	return
}
