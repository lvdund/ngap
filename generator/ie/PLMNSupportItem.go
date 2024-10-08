package ie

type PLMNSupportItem struct {
	PLMNIdentity     PLMNIdentity          `False,`
	SliceSupportList SliceSupportList      `False,`
	IEExtensions     PLMNSupportItemExtIEs `False,OPTIONAL`
}
