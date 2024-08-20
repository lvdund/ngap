package ie

type UeApplicationLayerMeasurementConfigurationInformation struct {
	QoeReference                                         []byte //`bitstring:"sizeLB:6,sizeUB:6"`
	ServiceType                                          *[]byte
	ChoiceAreaScopeOfQmc                                 *ChoiceAreaScopeOfQmc
	MeasurementCollectionEntityIpAddress                 *TransportLayerAddress
	QoeMeasurementStatus                                 *[]byte
	ContainerForApplicationLayerMeasurementConfiguration *[]byte
	MeasurementConfigurationApplicationLayerId           uint8 //`bitstring:"sizeLB:0,sizeUB:15"`
	SliceSupportListForQmc                               *SliceSupportListForQmc
	ChoiceMdtAlignmentInformation                        *ChoiceMdtAlignmentInformation
	AvailableRanVisibleQoeMetrics                        *AvailableRanVisibleQoeMetrics
}

type ChoiceAreaScopeOfQmc struct {
	CellBased     *CellBased
	TaBased       *TaBased
	TaiBased      *TaiBased
	PlmnAreaBased *PlmnAreaBased
}

type CellBased struct {
	CellIdListForQmc *CellIdListForQmc
}

type CellIdListForQmc struct {
	NgRanCgi *NgRanCgi
}

type TaBased struct {
	TaListForQmc *TaListForQmc
}

type TaListForQmc struct {
	Tac *Tac
}

type TaiBased struct {
	TaiListForQmc *TaiListForQmc
}

type TaiListForQmc struct {
	Tai *Tai
}

type PlmnAreaBased struct {
	PlmnListForQmc *PlmnListForQmc
}

type PlmnListForQmc struct {
	PlmnIdentity *PlmnIdentity
}

type SliceSupportListForQmc struct {
	SliceSupportQmcItem *SliceSupportQmcItem
}

type SliceSupportQmcItem struct {
	SNssai *SNssai
}

type ChoiceMdtAlignmentInformation struct {
	SBasedMdt *SBasedMdt
}

type SBasedMdt struct {
	NgRanTraceId []byte //`bitstring:"sizeLB:8,sizeUB:8"`
}
