package ie

type PwsFailureIndication struct {
	MessageType             MessageType             `bitstring:"sizeLB:0,sizeUB:150"`
	ChoicePwsFailedCellList ChoicePwsFailedCellList `bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeId         GlobalRanNodeId         `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoicePwsFailedCellList struct {
	EUtra EUtra `bitstring:"sizeLB:0,sizeUB:150"`
	Nr    Nr    `bitstring:"sizeLB:0,sizeUB:150"`
}

type PwsFailedEUtraCellList struct {
	EUtraCgi EUtraCgi `bitstring:"sizeLB:0,sizeUB:150"`
}

type PwsFailedNrCellList struct {
	NrCgi NrCgi `bitstring:"sizeLB:0,sizeUB:150"`
}
