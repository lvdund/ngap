package ie

type ServiceAreaInformationItem struct {
	PLMNIdentity   PLMNIdentity                     `False,`
	AllowedTACs    AllowedTACs                      `False,OPTIONAL`
	NotAllowedTACs NotAllowedTACs                   `False,OPTIONAL`
	IEExtensions   ServiceAreaInformationItemExtIEs `False,OPTIONAL`
}
