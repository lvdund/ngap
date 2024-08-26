package ie

import "ngap/aper"

type M1Configuration struct {
	M1ReportingTrigger                  []byte                              `bitstring:"sizeLB:0,sizeUB:150"`
	M1ThresholdEventA2                  M1ThresholdEventA2                  `bitstring:"sizeLB:0,sizeUB:150"`
	M1PeriodicReporting                 M1PeriodicReporting                 `bitstring:"sizeLB:0,sizeUB:150"`
	IncludeBeamMeasurementsIndication   []byte                              `bitstring:"sizeLB:0,sizeUB:150"`
	BeamMeasurementsReportConfiguration BeamMeasurementsReportConfiguration `bitstring:"sizeLB:0,sizeUB:150"`
}

type M1ThresholdEventA2 struct {
	ChoiceThresholdType ChoiceThresholdType `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceThresholdType struct {
	Rsrp Rsrp `bitstring:"sizeLB:0,sizeUB:150"`
	Rsrq Rsrq `bitstring:"sizeLB:0,sizeUB:150"`
	Sinr Sinr `bitstring:"sizeLB:0,sizeUB:150"`
}

type Rsrp struct {
	ThresholdRsrp aper.Integer `Integer:"valueLB:0,valueUB:127"`
}

type Rsrq struct {
	ThresholdRsrq aper.Integer `Integer:"valueLB:0,valueUB:127"`
}

type Sinr struct {
	ThresholdSinr aper.Integer `Integer:"valueLB:0,valueUB:127"`
}

type M1PeriodicReporting struct {
	ReportInterval         []byte `bitstring:"sizeLB:0,sizeUB:150"`
	ReportAmount           []byte `bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedReportInterval []byte `bitstring:"sizeLB:0,sizeUB:150"`
}

type BeamMeasurementsReportConfiguration struct {
	BeamMeasurementsReportQuantity BeamMeasurementsReportQuantity `bitstring:"sizeLB:0,sizeUB:150"`
	MaxnrofrsIndexestoReport       aper.Integer                   `Integer:"valueLB:1,valueUB:64"`
}

type BeamMeasurementsReportQuantity struct {
	Rsrp []byte `bitstring:"sizeLB:0,sizeUB:150"`
	Rsrq []byte `bitstring:"sizeLB:0,sizeUB:150"`
	Sinr []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
