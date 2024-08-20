package ie

type PwsFailureIndication struct {
	MessageType             *MessageType
	ChoicePwsFailedCellList *ChoicePwsFailedCellList
	GlobalRanNodeId         *GlobalRanNodeId
}

type ChoicePwsFailedCellList struct {
	EUtra *EUtra
	Nr    *Nr
}

type PwsFailedEUtraCellList struct {
	EUtraCgi *EUtraCgi
}

type PwsFailedNrCellList struct {
	NrCgi *NrCgi
}
