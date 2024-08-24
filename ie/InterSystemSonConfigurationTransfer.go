package ie

type InterSystemSonConfigurationTransfer struct {
ChoiceTransferType	ChoiceTransferType	//`bitstring:"sizeLB:0,sizeUB:150"`
InterSystemSonInformation	InterSystemSonInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceTransferType struct {
FromEUtranToNgRan	FromEUtranToNgRan	//`bitstring:"sizeLB:0,sizeUB:150"`
FromNgRanToEUtran	FromNgRanToEUtran	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type FromEUtranToNgRan struct {
SourceEnbId	SourceEnbId	//`bitstring:"sizeLB:0,sizeUB:150"`
TargetNgRanNodeId	TargetNgRanNodeId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type SourceEnbId struct {
GlobalEnbId	GlobalEnbId	//`bitstring:"sizeLB:0,sizeUB:150"`
SelectedEpsTai	EpsTai	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TargetNgRanNodeId struct {
GlobalRanNodeId	GlobalRanNodeId	//`bitstring:"sizeLB:0,sizeUB:150"`
SelectedTai	Tai	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type FromNgRanToEUtran struct {
SourceNgRanNodeId	SourceNgRanNodeId	//`bitstring:"sizeLB:0,sizeUB:150"`
TargetEnbId	TargetEnbId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type SourceNgRanNodeId struct {
GlobalRanNodeId	GlobalRanNodeId	//`bitstring:"sizeLB:0,sizeUB:150"`
SelectedTai	Tai	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TargetEnbId struct {
GlobalEnbId	GlobalEnbId	//`bitstring:"sizeLB:0,sizeUB:150"`
SelectedEpsTai	EpsTai	//`bitstring:"sizeLB:0,sizeUB:150"`
}
