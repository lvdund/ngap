package ie

import "ngap/aper"

type RerouteNasRequest struct {
	MessageType                         MessageType                         //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                         RanUeNgapId                         //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                         AmfUeNgapId                         //`bitstring:"sizeLB:0,sizeUB:150"`
	NgapMessage                         aper.OctetString                    //`octetstring:"sizeLB:0,sizeUB:150"`
	AmfSetId                            AmfSetId                            //`bitstring:"sizeLB:0,sizeUB:150"`
	AllowedNssai                        AllowedNssai                        //`bitstring:"sizeLB:0,sizeUB:150"`
	SourceToTargetAmfInformationReroute SourceToTargetAmfInformationReroute //`bitstring:"sizeLB:0,sizeUB:150"`
}
