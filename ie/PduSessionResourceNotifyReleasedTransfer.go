package ie

type PduSessionResourceNotifyReleasedTransfer struct {
	Cause                        Cause                        //`bitstring:"sizeLB:0,sizeUB:150"`
	SecondaryRatUsageInformation SecondaryRatUsageInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}
