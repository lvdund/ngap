package ie

type SONInformationReply struct {
	XnTNLConfigurationInfo XnTNLConfigurationInfo    `True,OPTIONAL`
	IEExtensions           SONInformationReplyExtIEs `False,OPTIONAL`
}
