package ie

type GlobalRANNodeID struct {
	GlobalGNBID      GlobalGNBID           `True,`
	GlobalNgENBID    GlobalNgENBID         `True,`
	GlobalN3IWFID    GlobalN3IWFID         `True,`
	ChoiceExtensions GlobalRANNodeIDExtIEs `False,`
}
