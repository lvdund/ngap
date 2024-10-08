package ie

type FiveGSTMSI struct {
	AMFSetID     AMFSetID         `False,`
	AMFPointer   AMFPointer       `False,`
	FiveGTMSI    FiveGTMSI        `False,`
	IEExtensions FiveGSTMSIExtIEs `False,OPTIONAL`
}
