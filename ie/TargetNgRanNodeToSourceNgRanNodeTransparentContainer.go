package ie

import "ngap/aper"

type TargetNgRanNodeToSourceNgRanNodeTransparentContainer struct {
RrcContainer	aper.OctetString	//`octetstring:"sizeLB:0,sizeUB:150"`
DapsResponseInformationList	[]DapsResponseInformationItem	//`bitstring:"sizeLB:0,sizeUB:150"`
DirectForwardingPathAvailability	DirectForwardingPathAvailability	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsActiveSessionInformationTargetToSourceList	[]MbsActiveSessionInformationTargetToSourceItem	//`bitstring:"sizeLB:0,sizeUB:150"`
NgapIeSupportInformationResponseList	NgapIeSupportInformationResponseList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type DapsResponseInformationItem struct {
DrbId	DrbId	//`bitstring:"sizeLB:0,sizeUB:150"`
DapsResponseInformation	DapsResponseInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsActiveSessionInformationTargetToSourceItem struct {
MbsSessionId	MbsSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
DataForwardingResponseMrbList	[]DataForwardingResponseMrbItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type DataForwardingResponseMrbItem struct {
MrbId	MrbId	//`bitstring:"sizeLB:0,sizeUB:150"`
DlForwardingUpTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
MrbProgressInformation	MrbProgressInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}
