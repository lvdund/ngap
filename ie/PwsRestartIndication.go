package ie

type PwsRestartIndication struct {
	MessageType                   *MessageType
	ChoiceCellListForRestart      *ChoiceCellListForRestart
	GlobalRanNodeId               *GlobalRanNodeId
	TaiListForRestart             *TaiListForRestart
	EmergencyAreaIdListForRestart *EmergencyAreaIdListForRestart
}

type ChoiceCellListForRestart struct {
	EUtra *EUtra
	Nr    *Nr
}

type EUtra struct {
	EUtraCellListForRestart *EUtraCellListForRestart
}

type EUtraCellListForRestart struct {
	EUtraCgi *EUtraCgi
}

type Nr struct {
	NrCellListForRestart *NrCellListForRestart
}

type NrCellListForRestart struct {
	NrCgi *NrCgi
}

type TaiListForRestart struct {
	Tai *Tai
}

type EmergencyAreaIdListForRestart struct {
	EmergencyAreaId *EmergencyAreaId
}
