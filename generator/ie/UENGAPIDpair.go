package ie

type UENGAPIDpair struct {
	AMFUENGAPID  AMFUENGAPID        `False,`
	RANUENGAPID  RANUENGAPID        `False,`
	IEExtensions UENGAPIDpairExtIEs `False,OPTIONAL`
}
