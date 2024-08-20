package ie

type WlanMeasurementConfiguration struct {
WlanMeasurementConfiguration	*[]byte
WlanMeasurementConfigurationNameList	*[]WlanMeasurementConfigurationNameItem
WlanRssi	*[]byte
WlanRtt	*[]byte
}

type WlanMeasurementConfigurationNameItem struct {
WlanMeasurementConfigurationName	*[]byte
}
