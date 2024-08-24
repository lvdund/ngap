package ie

import "ngap/aper"

type UeApplicationLayerMeasurementConfigurationInformation struct {
	QoeReference                                         aper.OctetString              //`octetstring:"sizeLB:6,sizeUB:6"`
	ServiceType                                          []byte                        //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceAreaScopeOfQmc                                 ChoiceAreaScopeOfQmc          //`bitstring:"sizeLB:0,sizeUB:150"`
	MeasurementCollectionEntityIpAddress                 TransportLayerAddress         //`bitstring:"sizeLB:0,sizeUB:150"`
	QoeMeasurementStatus                                 []byte                        //`bitstring:"sizeLB:0,sizeUB:150"`
	ContainerForApplicationLayerMeasurementConfiguration aper.OctetString              //`octetstring:"sizeLB:0,sizeUB:150"`
	MeasurementConfigurationApplicationLayerId           aper.Integer                  //`Integer:"valueLB:0,valueUB:15"`
	SliceSupportListForQmc                               SliceSupportListForQmc        //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceMdtAlignmentInformation                        ChoiceMdtAlignmentInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	AvailableRanVisibleQoeMetrics                        AvailableRanVisibleQoeMetrics //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceAreaScopeOfQmc struct {
	CellBased     CellBased     //`bitstring:"sizeLB:0,sizeUB:150"`
	TaBased       TaBased       //`bitstring:"sizeLB:0,sizeUB:150"`
	TaiBased      TaiBased      //`bitstring:"sizeLB:0,sizeUB:150"`
	PlmnAreaBased PlmnAreaBased //`bitstring:"sizeLB:0,sizeUB:150"`
}

type CellBased struct {
	CellIdListForQmc CellIdListForQmc //`bitstring:"sizeLB:0,sizeUB:150"`
}

type CellIdListForQmc struct {
	NgRanCgi NgRanCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaBased struct {
	TaListForQmc TaListForQmc //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaListForQmc struct {
	Tac Tac //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiBased struct {
	TaiListForQmc TaiListForQmc //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiListForQmc struct {
	Tai Tai //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PlmnAreaBased struct {
	PlmnListForQmc PlmnListForQmc //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PlmnListForQmc struct {
	PlmnIdentity PlmnIdentity //`bitstring:"sizeLB:0,sizeUB:150"`
}

type SliceSupportListForQmc struct {
	SliceSupportQmcItem SliceSupportQmcItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type SliceSupportQmcItem struct {
	SNssai SNssai //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceMdtAlignmentInformation struct {
	SBasedMdt SBasedMdt //`bitstring:"sizeLB:0,sizeUB:150"`
}

type SBasedMdt struct {
	NgRanTraceId aper.OctetString //`octetstring:"sizeLB:8,sizeUB:8"`
}
