package ie

import "ngap/aper"

type NgapProtocolIeId struct {
	NgapProtocolIeId aper.Integer `Integer:"valueLB:0,valueUB:65535"`
}

func (ie *NgapProtocolIeId) Decode(r aper.AperReader) error {

	return nil
}

func (ie *NgapProtocolIeId) Encode(r aper.AperWriter) (err error) {
	if err = r.WriteInteger(ie.NgapProtocolIeId, &aper.Constrain{Lb: 10, Ub: 10}, true); err != nil {
		return err
	}
	return nil
}
