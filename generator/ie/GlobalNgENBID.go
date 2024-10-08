package ie

type GlobalNgENBID struct {
	PLMNIdentity PLMNIdentity        `False,`
	NgENBID      NgENBID             `False,`
	IEExtensions GlobalNgENBIDExtIEs `False,OPTIONAL`
}
