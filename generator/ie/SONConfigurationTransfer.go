package ie

type SONConfigurationTransfer struct {
	TargetRANNodeID        TargetRANNodeID                `True,`
	SourceRANNodeID        SourceRANNodeID                `True,`
	SONInformation         SONInformation                 `False,`
	XnTNLConfigurationInfo XnTNLConfigurationInfo         `True,OPTIONAL`
	IEExtensions           SONConfigurationTransferExtIEs `False,OPTIONAL`
}
