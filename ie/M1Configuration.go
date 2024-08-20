package ie

type M1Configuration struct {
	M1ReportingTrigger                  *[]byte
	M1ThresholdEventA2                  *M1ThresholdEventA2
	M1PeriodicReporting                 *M1PeriodicReporting
	IncludeBeamMeasurementsIndication   *[]byte
	BeamMeasurementsReportConfiguration *BeamMeasurementsReportConfiguration
}

type M1ThresholdEventA2 struct {
	ChoiceThresholdType *ChoiceThresholdType
}

type ChoiceThresholdType struct {
	Rsrp *Rsrp
	Rsrq *Rsrq
	Sinr *Sinr
}

type Rsrp struct {
	ThresholdRsrp uint8 //`bitstring:"sizeLB:0,sizeUB:127"`
}

type Rsrq struct {
	ThresholdRsrq uint8 //`bitstring:"sizeLB:0,sizeUB:127"`
}

type Sinr struct {
	ThresholdSinr uint8 //`bitstring:"sizeLB:0,sizeUB:127"`
}

type M1PeriodicReporting struct {
	ReportInterval         *[]byte
	ReportAmount           *[]byte
	ExtendedReportInterval *[]byte
}

type BeamMeasurementsReportConfiguration struct {
	BeamMeasurementsReportQuantity *BeamMeasurementsReportQuantity
	MaxnrofrsIndexestoReport       uint8 //`bitstring:"sizeLB:1,sizeUB:64"`
}

type BeamMeasurementsReportQuantity struct {
	Rsrp *[]byte
	Rsrq *[]byte
	Sinr *[]byte
}
