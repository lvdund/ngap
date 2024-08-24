package ie

type WarningAreaList struct {
	ChoiceWarningArea ChoiceWarningArea //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceWarningArea struct {
	EUtraCellIds     EUtraCellIds     //`bitstring:"sizeLB:0,sizeUB:150"`
	NrCellIds        NrCellIds        //`bitstring:"sizeLB:0,sizeUB:150"`
	TaisForWarning   TaisForWarning   //`bitstring:"sizeLB:0,sizeUB:150"`
	EmergencyAreaIds EmergencyAreaIds //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EUtraCellIds struct {
	EutraCgiListForWarning EutraCgiListForWarning //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EutraCgiListForWarning struct {
	EUtraCgi EUtraCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NrCellIds struct {
	NrCgiListForWarning NrCgiListForWarning //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NrCgiListForWarning struct {
	NrCgi NrCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaisForWarning struct {
	TaiListForWarning TaiListForWarning //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiListForWarning struct {
	Tai Tai //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIds struct {
	EmergencyAreaIdList EmergencyAreaIdList //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EmergencyAreaIdList struct {
	EmergencyAreaId EmergencyAreaId //`bitstring:"sizeLB:0,sizeUB:150"`
}
