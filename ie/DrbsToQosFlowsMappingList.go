package ie

type DrbsToQosFlowsMappingList struct {
	DrbsToQosFlowsMappingItem *DrbsToQosFlowsMappingItem
}

type DrbsToQosFlowsMappingItem struct {
	DrbId                  *DrbId
	AssociatedQosFlowList  *AssociatedQosFlowList
	DapsRequestInformation *DapsRequestInformation
}
