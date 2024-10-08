package ie

type HandoverRequiredTransfer struct {
	DirectForwardingPathAvailability DirectForwardingPathAvailability `False,OPTIONAL`
	IEExtensions                     HandoverRequiredTransferExtIEs   `False,OPTIONAL`
}
