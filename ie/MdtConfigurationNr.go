package ie

type MdtConfigurationNr struct {
	MdtActivation              *[]byte
	ChoiceAreaScopeOfMdt       *ChoiceAreaScopeOfMdt
	ChoiceMdtMode              *ChoiceMdtMode
	SignallingBasedMdtPlmnList *MdtPlmnList
}

type ChoiceAreaScopeOfMdt struct {
	CellBased *CellBased
	TaBased   *TaBased
	PlmnWide  *[]byte
	TaiBased  *TaiBased
}

type CellIdListForMdt struct {
	NrCgi *NrCgi
}

type TaListForMdt struct {
	Tac *Tac
}

type TaiListForMdt struct {
	Tai *Tai
}

type ChoiceMdtMode struct {
	ImmediateMdt *ImmediateMdt
	LoggedMdt    *LoggedMdt
}

type ImmediateMdt struct {
	MeasurementsToActivate            []byte //`bitstring:"sizeLB:8,sizeUB:8"`
	M1Configuration                   *M1Configuration
	M4Configuration                   *M4Configuration
	M5Configuration                   *M5Configuration
	M6Configuration                   *M6Configuration
	M7Configuration                   *M7Configuration
	BluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration
	WlanMeasurementConfiguration      *WlanMeasurementConfiguration
	MdtLocationInformation            *MdtLocationInformation
	SensorMeasurementConfiguration    *SensorMeasurementConfiguration
}

type LoggedMdt struct {
	LoggingInterval                   *[]byte
	LoggingDuration                   *[]byte
	ChoiceReportType                  *ChoiceReportType
	BluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration
	WlanMeasurementConfiguration      *WlanMeasurementConfiguration
	SensorMeasurementConfiguration    *SensorMeasurementConfiguration
	AreaScopeOfNeighbourCells         *AreaScopeOfNeighbourCells
	EarlyMeasurement                  *[]byte
}

type EventTriggered struct {
	EventTriggerLoggedMdtConfiguration *EventTriggerLoggedMdtConfiguration
}
