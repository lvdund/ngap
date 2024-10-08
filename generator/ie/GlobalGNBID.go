package ie

type GlobalGNBID struct {
	PLMNIdentity PLMNIdentity      `False,`
	GNBID        GNBID             `False,`
	IEExtensions GlobalGNBIDExtIEs `False,OPTIONAL`
}
