package ie

type RanStatusTransferTransparentContainer struct {
	DrbsSubjectToStatusTransferList []DrbsSubjectToStatusTransferItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type DrbsSubjectToStatusTransferItem struct {
	DrbId                                       DrbId                 `bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceUlDrbStatus                           ChoiceUlDrbStatus     `bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceDlDrbStatus                           ChoiceDlDrbStatus     `bitstring:"sizeLB:0,sizeUB:150"`
	OldAssociatedQosFlowListUlEndMarkerExpected AssociatedQosFlowList `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceUlDrbStatus struct {
	Ie12Bits Ie12Bits `bitstring:"sizeLB:0,sizeUB:150"`
	Ie18Bits Ie18Bits `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceDlDrbStatus struct {
	Ie12Bits Ie12Bits `bitstring:"sizeLB:0,sizeUB:150"`
	Ie18Bits Ie18Bits `bitstring:"sizeLB:0,sizeUB:150"`
}
