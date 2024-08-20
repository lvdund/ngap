package ie

type EarlyStatusTransferTransparentContainer struct {
	ChoiceProcedureStage *ChoiceProcedureStage
}

type ChoiceProcedureStage struct {
	FirstDlCount *FirstDlCount
}

type FirstDlCount struct {
	DrbsSubjectToEarlyStatusTransferList *[]DrbsSubjectToEarlyStatusTransferItem
}

type DrbsSubjectToEarlyStatusTransferItem struct {
	DrbId              *DrbId
	ChoiceFirstDlCount *ChoiceFirstDlCount
}

type ChoiceFirstDlCount struct {
	Ie12Bits *Ie12Bits
	Ie18Bits *Ie18Bits
}

type Ie12Bits struct {
	FirstDlCountValue *CountValueForPdcpSnLength12
}

type Ie18Bits struct {
	FirstDlCountValue *CountValueForPdcpSnLength18
}
