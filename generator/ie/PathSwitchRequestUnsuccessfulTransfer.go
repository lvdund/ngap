package ie

type PathSwitchRequestUnsuccessfulTransfer struct {
	Cause        Cause                                       `False,`
	IEExtensions PathSwitchRequestUnsuccessfulTransferExtIEs `False,OPTIONAL`
}
