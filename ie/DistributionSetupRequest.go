package ie

import "ngap/aper"

type DistributionSetupRequest struct {
	MessageType                         MessageType      //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId                        MbsSessionId     //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId                    MbsAreaSessionId //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsDistributionSetupRequestTransfer aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}
