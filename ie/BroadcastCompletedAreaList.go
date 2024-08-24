package ie

type BroadcastCompletedAreaList struct {
ChoiceBroadcastCompletedArea	ChoiceBroadcastCompletedArea	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceBroadcastCompletedArea struct {
CellIdBroadcastEUtra	CellIdBroadcastEUtra	//`bitstring:"sizeLB:0,sizeUB:150"`
TaiBroadcastEUtra	TaiBroadcastEUtra	//`bitstring:"sizeLB:0,sizeUB:150"`
EmergencyAreaIdBroadcastEUtra	EmergencyAreaIdBroadcastEUtra	//`bitstring:"sizeLB:0,sizeUB:150"`
CellIdBroadcastNr	CellIdBroadcastNr	//`bitstring:"sizeLB:0,sizeUB:150"`
TaiBroadcastNr	TaiBroadcastNr	//`bitstring:"sizeLB:0,sizeUB:150"`
EmergencyAreaIdBroadcastNr	EmergencyAreaIdBroadcastNr	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type CellIdBroadcastEUtra struct {
CompletedCellList	CompletedCellList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type CompletedCellList struct {
NrCgi	NrCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiBroadcastEUtra struct {
TaiBroadcast	TaiBroadcast	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiBroadcast struct {
Tai	Tai	//`bitstring:"sizeLB:0,sizeUB:150"`
CompletedCellInTaiList	CompletedCellInTaiList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type CompletedCellInTaiList struct {
NrCgi	NrCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIdBroadcastEUtra struct {
EmergencyAreaIdBroadcast	EmergencyAreaIdBroadcast	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIdBroadcast struct {
EmergencyAreaId	EmergencyAreaId	//`bitstring:"sizeLB:0,sizeUB:150"`
CompletedCellInEmergencyAreaIdList	CompletedCellInEmergencyAreaIdList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type CompletedCellInEmergencyAreaIdList struct {
NrCgi	NrCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type CellIdBroadcastNr struct {
CompletedCellList	CompletedCellList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiBroadcastNr struct {
TaiBroadcast	TaiBroadcast	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIdBroadcastNr struct {
EmergencyAreaIdBroadcast	EmergencyAreaIdBroadcast	//`bitstring:"sizeLB:0,sizeUB:150"`
}
