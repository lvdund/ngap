package ie

type ForbiddenAreaInformationItem struct {
	PLMNIdentity  PLMNIdentity                       `False,`
	ForbiddenTACs ForbiddenTACs                      `False,`
	IEExtensions  ForbiddenAreaInformationItemExtIEs `False,OPTIONAL`
}
