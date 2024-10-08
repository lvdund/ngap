package ie

import "ngap/aper"

type TraceActivation struct {
	NgRanTraceId                   aper.OctetString      `octetstring:"sizeLB:8,sizeUB:8"`
	InterfacesToTrace              aper.BitString        `bitstring:"sizeLB:8,sizeUB:8"`
	TraceDepth                     []byte                `bitstring:"sizeLB:0,sizeUB:150"`
	TraceCollectionEntityIpAddress TransportLayerAddress `bitstring:"sizeLB:0,sizeUB:150"`
	MdtConfiguration               MdtConfiguration      `bitstring:"sizeLB:0,sizeUB:150"`
	TraceCollectionEntityUri       Uri                   `bitstring:"sizeLB:0,sizeUB:150"`
}
