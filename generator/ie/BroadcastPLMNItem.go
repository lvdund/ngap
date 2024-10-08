package ie

type BroadcastPLMNItem struct {
	PLMNIdentity        PLMNIdentity            `False,`
	TAISliceSupportList SliceSupportList        `False,`
	IEExtensions        BroadcastPLMNItemExtIEs `False,OPTIONAL`
}
