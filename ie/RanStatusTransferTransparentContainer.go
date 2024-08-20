package ie

type RanStatusTransferTransparentContainer struct {
DrbsSubjectToStatusTransferList	*[]DrbsSubjectToStatusTransferItem
}

type DrbsSubjectToStatusTransferItem struct {
DrbId	*DrbId
ChoiceUlDrbStatus	*ChoiceUlDrbStatus
ChoiceDlDrbStatus	*ChoiceDlDrbStatus
OldAssociatedQosFlowListUlEndMarkerExpected	*AssociatedQosFlowList
}

type ChoiceUlDrbStatus struct {
Ie12Bits	*Ie12Bits
Ie18Bits	*Ie18Bits
}

type ChoiceDlDrbStatus struct {
Ie12Bits	*Ie12Bits
Ie18Bits	*Ie18Bits
}
