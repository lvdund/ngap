package ie

type SensorMeasurementConfiguration struct {
	SensorMeasurementConfiguration         []byte                                   `bitstring:"sizeLB:0,sizeUB:150"`
	SensorMeasurementConfigurationNameList []SensorMeasurementConfigurationNameItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type SensorMeasurementConfigurationNameItem struct {
	ChoiceSensorName ChoiceSensorName `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceSensorName struct {
	UncompensatedBarometric UncompensatedBarometric `bitstring:"sizeLB:0,sizeUB:150"`
	UeSpeed                 UeSpeed                 `bitstring:"sizeLB:0,sizeUB:150"`
	UeOrientation           UeOrientation           `bitstring:"sizeLB:0,sizeUB:150"`
}

type UncompensatedBarometric struct {
	UncompensatedBarometricConfiguration []byte `bitstring:"sizeLB:0,sizeUB:150"`
}

type UeSpeed struct {
	UeSpeedConfiguration []byte `bitstring:"sizeLB:0,sizeUB:150"`
}

type UeOrientation struct {
	UeOrientationConfiguration []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
