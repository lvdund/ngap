package ie

type PwsRestartIndication struct {
	MessageType                   MessageType                   `bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceCellListForRestart      ChoiceCellListForRestart      `bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeId               GlobalRanNodeId               `bitstring:"sizeLB:0,sizeUB:150"`
	TaiListForRestart             TaiListForRestart             `bitstring:"sizeLB:0,sizeUB:150"`
	EmergencyAreaIdListForRestart EmergencyAreaIdListForRestart `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceCellListForRestart struct {
	EUtra EUtra `bitstring:"sizeLB:0,sizeUB:150"`
	Nr    Nr    `bitstring:"sizeLB:0,sizeUB:150"`
}

type EUtra struct {
	EUtraCellListForRestart EUtraCellListForRestart `bitstring:"sizeLB:0,sizeUB:150"`
}

type EUtraCellListForRestart struct {
	EUtraCgi EUtraCgi `bitstring:"sizeLB:0,sizeUB:150"`
}

type Nr struct {
	NrCellListForRestart NrCellListForRestart `bitstring:"sizeLB:0,sizeUB:150"`
}

type NrCellListForRestart struct {
	NrCgi NrCgi `bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiListForRestart struct {
	Tai Tai `bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIdListForRestart struct {
	EmergencyAreaId EmergencyAreaId `bitstring:"sizeLB:0,sizeUB:150"`
}
