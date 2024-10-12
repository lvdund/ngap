package ie

type PWSRestartIndicationIEs struct {
	CellIDListForRestart          CellIDListForRestart          `,reject,mandatory`
	GlobalRANNodeID               GlobalRANNodeID               `,reject,mandatory`
	TAIListForRestart             TAIListForRestart             `,reject,mandatory`
	EmergencyAreaIDListForRestart EmergencyAreaIDListForRestart `,reject,optional`
}
