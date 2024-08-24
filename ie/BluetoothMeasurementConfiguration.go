package ie

import "ngap/aper"

type BluetoothMeasurementConfiguration struct {
BluetoothMeasurementConfiguration	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
BluetoothMeasurementConfigurationNameList	[]BluetoothMeasurementConfigurationNameItem	//`bitstring:"sizeLB:0,sizeUB:150"`
BtRssi	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type BluetoothMeasurementConfigurationNameItem struct {
BluetoothMeasurementConfigurationName	aper.OctetString	//`octetstring:"sizeLB:0,sizeUB:150"`
}
