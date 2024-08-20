package ie

type WarningAreaList struct {
ChoiceWarningArea	*ChoiceWarningArea
}

type ChoiceWarningArea struct {
EUtraCellIds	*EUtraCellIds
NrCellIds	*NrCellIds
TaisForWarning	*TaisForWarning
EmergencyAreaIds	*EmergencyAreaIds
}

type EUtraCellIds struct {
EutraCgiListForWarning	*EutraCgiListForWarning
}

type EutraCgiListForWarning struct {
EUtraCgi	*EUtraCgi
}

type NrCellIds struct {
NrCgiListForWarning	*NrCgiListForWarning
}

type NrCgiListForWarning struct {
NrCgi	*NrCgi
}

type TaisForWarning struct {
TaiListForWarning	*TaiListForWarning
}

type TaiListForWarning struct {
Tai	*Tai
}

type EmergencyAreaIds struct {
EmergencyAreaIdList	*EmergencyAreaIdList
}

type EmergencyAreaIdList struct {
EmergencyAreaId	*EmergencyAreaId
}
