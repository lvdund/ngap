package ie

type UserLocationInformationEUTRA struct {
	EUTRACGI     EUTRACGI                           `True,`
	TAI          TAI                                `True,`
	TimeStamp    TimeStamp                          `False,OPTIONAL`
	IEExtensions UserLocationInformationEUTRAExtIEs `False,OPTIONAL`
}
