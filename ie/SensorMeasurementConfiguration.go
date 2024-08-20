package ie

type SensorMeasurementConfiguration struct {
SensorMeasurementConfiguration	*[]byte
SensorMeasurementConfigurationNameList	*[]SensorMeasurementConfigurationNameItem
}

type SensorMeasurementConfigurationNameItem struct {
ChoiceSensorName	*ChoiceSensorName
}

type ChoiceSensorName struct {
UncompensatedBarometric	*UncompensatedBarometric
UeSpeed	*UeSpeed
UeOrientation	*UeOrientation
}

type UncompensatedBarometric struct {
UncompensatedBarometricConfiguration	*[]byte
}

type UeSpeed struct {
UeSpeedConfiguration	*[]byte
}

type UeOrientation struct {
UeOrientationConfiguration	*[]byte
}
