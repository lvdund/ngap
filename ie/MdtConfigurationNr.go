package ie

type MdtConfigurationNr struct {
	MdtActivation              []byte               //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceAreaScopeOfMdt       ChoiceAreaScopeOfMdt //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceMdtMode              ChoiceMdtMode        //`bitstring:"sizeLB:0,sizeUB:150"`
	SignallingBasedMdtPlmnList MdtPlmnList          //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceAreaScopeOfMdt struct {
	CellBased CellBased //`bitstring:"sizeLB:0,sizeUB:150"`
	TaBased   TaBased   //`bitstring:"sizeLB:0,sizeUB:150"`
	PlmnWide  []byte    //`bitstring:"sizeLB:0,sizeUB:150"`
	TaiBased  TaiBased  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type CellIdListForMdt struct {
	NrCgi NrCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaListForMdt struct {
	Tac Tac //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiListForMdt struct {
	Tai Tai //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceMdtMode struct {
	ImmediateMdt ImmediateMdt //`bitstring:"sizeLB:0,sizeUB:150"`
	LoggedMdt    LoggedMdt    //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ImmediateMdt struct {
	MeasurementsToActivate            []byte                            //`bitstring:"sizeLB:8,sizeUB:8"`
	M1Configuration                   M1Configuration                   //`bitstring:"sizeLB:0,sizeUB:150"`
	M4Configuration                   M4Configuration                   //`bitstring:"sizeLB:0,sizeUB:150"`
	M5Configuration                   M5Configuration                   //`bitstring:"sizeLB:0,sizeUB:150"`
	M6Configuration                   M6Configuration                   //`bitstring:"sizeLB:0,sizeUB:150"`
	M7Configuration                   M7Configuration                   //`bitstring:"sizeLB:0,sizeUB:150"`
	BluetoothMeasurementConfiguration BluetoothMeasurementConfiguration //`bitstring:"sizeLB:0,sizeUB:150"`
	WlanMeasurementConfiguration      WlanMeasurementConfiguration      //`bitstring:"sizeLB:0,sizeUB:150"`
	MdtLocationInformation            MdtLocationInformation            //`bitstring:"sizeLB:0,sizeUB:150"`
	SensorMeasurementConfiguration    SensorMeasurementConfiguration    //`bitstring:"sizeLB:0,sizeUB:150"`
}

type LoggedMdt struct {
	LoggingInterval                   []byte                            //`bitstring:"sizeLB:0,sizeUB:150"`
	LoggingDuration                   []byte                            //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceReportType                  ChoiceReportType                  //`bitstring:"sizeLB:0,sizeUB:150"`
	BluetoothMeasurementConfiguration BluetoothMeasurementConfiguration //`bitstring:"sizeLB:0,sizeUB:150"`
	WlanMeasurementConfiguration      WlanMeasurementConfiguration      //`bitstring:"sizeLB:0,sizeUB:150"`
	SensorMeasurementConfiguration    SensorMeasurementConfiguration    //`bitstring:"sizeLB:0,sizeUB:150"`
	AreaScopeOfNeighbourCells         AreaScopeOfNeighbourCells         //`bitstring:"sizeLB:0,sizeUB:150"`
	EarlyMeasurement                  []byte                            //`bitstring:"sizeLB:0,sizeUB:150"`
}

type EventTriggered struct {
	EventTriggerLoggedMdtConfiguration EventTriggerLoggedMdtConfiguration //`bitstring:"sizeLB:0,sizeUB:150"`
}
