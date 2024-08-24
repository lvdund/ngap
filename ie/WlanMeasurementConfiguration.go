package ie

import "ngap/aper"

type WlanMeasurementConfiguration struct {
	WlanMeasurementConfiguration         []byte                                 //`bitstring:"sizeLB:0,sizeUB:150"`
	WlanMeasurementConfigurationNameList []WlanMeasurementConfigurationNameItem //`bitstring:"sizeLB:0,sizeUB:150"`
	WlanRssi                             []byte                                 //`bitstring:"sizeLB:0,sizeUB:150"`
	WlanRtt                              []byte                                 //`bitstring:"sizeLB:0,sizeUB:150"`
}

type WlanMeasurementConfigurationNameItem struct {
	WlanMeasurementConfigurationName aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}
