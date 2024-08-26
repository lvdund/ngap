package ie

type DrbsToQosFlowsMappingList struct {
	DrbsToQosFlowsMappingItem DrbsToQosFlowsMappingItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type DrbsToQosFlowsMappingItem struct {
	DrbId                  DrbId                  `bitstring:"sizeLB:0,sizeUB:150"`
	AssociatedQosFlowList  AssociatedQosFlowList  `bitstring:"sizeLB:0,sizeUB:150"`
	DapsRequestInformation DapsRequestInformation `bitstring:"sizeLB:0,sizeUB:150"`
}
