package ie

type BroadcastCompletedAreaList struct {
	ChoiceBroadcastCompletedArea *ChoiceBroadcastCompletedArea
}

type ChoiceBroadcastCompletedArea struct {
	CellIdBroadcastEUtra          *CellIdBroadcastEUtra
	TaiBroadcastEUtra             *TaiBroadcastEUtra
	EmergencyAreaIdBroadcastEUtra *EmergencyAreaIdBroadcastEUtra
	CellIdBroadcastNr             *CellIdBroadcastNr
	TaiBroadcastNr                *TaiBroadcastNr
	EmergencyAreaIdBroadcastNr    *EmergencyAreaIdBroadcastNr
}

type CellIdBroadcastEUtra struct {
	CompletedCellList *CompletedCellList
}

type CompletedCellList struct {
	NrCgi *NrCgi
}

type TaiBroadcastEUtra struct {
	TaiBroadcast *TaiBroadcast
}

type TaiBroadcast struct {
	Tai                    *Tai
	CompletedCellInTaiList *CompletedCellInTaiList
}

type CompletedCellInTaiList struct {
	NrCgi *NrCgi
}

type EmergencyAreaIdBroadcastEUtra struct {
	EmergencyAreaIdBroadcast *EmergencyAreaIdBroadcast
}

type EmergencyAreaIdBroadcast struct {
	EmergencyAreaId                    *EmergencyAreaId
	CompletedCellInEmergencyAreaIdList *CompletedCellInEmergencyAreaIdList
}

type CompletedCellInEmergencyAreaIdList struct {
	NrCgi *NrCgi
}

type CellIdBroadcastNr struct {
	CompletedCellList *CompletedCellList
}

type TaiBroadcastNr struct {
	TaiBroadcast *TaiBroadcast
}

type EmergencyAreaIdBroadcastNr struct {
	EmergencyAreaIdBroadcast *EmergencyAreaIdBroadcast
}
