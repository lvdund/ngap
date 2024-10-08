package ie

type ERABInformationItem struct {
	ERABID       ERABID                    `False,`
	DLForwarding DLForwarding              `False,OPTIONAL`
	IEExtensions ERABInformationItemExtIEs `False,OPTIONAL`
}
