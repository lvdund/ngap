package ie

type RimInformation struct {
TargetGnbSetId	GnbSetId	//`bitstring:"sizeLB:0,sizeUB:150"`
RimRsDetection	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
