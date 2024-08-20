package ie

type BluetoothMeasurementConfiguration struct {
BluetoothMeasurementConfiguration	*[]byte
BluetoothMeasurementConfigurationNameList	*[]BluetoothMeasurementConfigurationNameItem
BtRssi	*[]byte
}

type BluetoothMeasurementConfigurationNameItem struct {
BluetoothMeasurementConfigurationName	*[]byte
}
