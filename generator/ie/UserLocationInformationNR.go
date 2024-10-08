package ie

type UserLocationInformationNR struct {
	NRCGI        NRCGI                           `True,`
	TAI          TAI                             `True,`
	TimeStamp    TimeStamp                       `False,OPTIONAL`
	IEExtensions UserLocationInformationNRExtIEs `False,OPTIONAL`
}
