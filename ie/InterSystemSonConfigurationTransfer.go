package ie

type InterSystemSonConfigurationTransfer struct {
	ChoiceTransferType        *ChoiceTransferType
	InterSystemSonInformation *InterSystemSonInformation
}

type ChoiceTransferType struct {
	FromEUtranToNgRan *FromEUtranToNgRan
	FromNgRanToEUtran *FromNgRanToEUtran
}

type FromEUtranToNgRan struct {
	SourceEnbId       *SourceEnbId
	TargetNgRanNodeId *TargetNgRanNodeId
}

type SourceEnbId struct {
	GlobalEnbId    *GlobalEnbId
	SelectedEpsTai *EpsTai
}

type TargetNgRanNodeId struct {
	GlobalRanNodeId *GlobalRanNodeId
	SelectedTai     *Tai
}

type FromNgRanToEUtran struct {
	SourceNgRanNodeId *SourceNgRanNodeId
	TargetEnbId       *TargetEnbId
}

type SourceNgRanNodeId struct {
	GlobalRanNodeId *GlobalRanNodeId
	SelectedTai     *Tai
}

type TargetEnbId struct {
	GlobalEnbId    *GlobalEnbId
	SelectedEpsTai *EpsTai
}
