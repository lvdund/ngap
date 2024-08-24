package ie

type EarlyStatusTransferTransparentContainer struct {
	ChoiceProcedureStage ChoiceProcedureStage //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceProcedureStage struct {
	FirstDlCount FirstDlCount //`bitstring:"sizeLB:0,sizeUB:150"`
}

type FirstDlCount struct {
	DrbsSubjectToEarlyStatusTransferList []DrbsSubjectToEarlyStatusTransferItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type DrbsSubjectToEarlyStatusTransferItem struct {
	DrbId              DrbId              //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceFirstDlCount ChoiceFirstDlCount //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceFirstDlCount struct {
	Ie12Bits Ie12Bits //`bitstring:"sizeLB:0,sizeUB:150"`
	Ie18Bits Ie18Bits //`bitstring:"sizeLB:0,sizeUB:150"`
}

type Ie12Bits struct {
	FirstDlCountValue CountValueForPdcpSnLength12 //`bitstring:"sizeLB:0,sizeUB:150"`
}

type Ie18Bits struct {
	FirstDlCountValue CountValueForPdcpSnLength18 //`bitstring:"sizeLB:0,sizeUB:150"`
}
