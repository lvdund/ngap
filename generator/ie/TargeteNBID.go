package ie

type TargeteNBID struct {
	GlobalENBID    GlobalNgENBID     `True,`
	SelectedEPSTAI EPSTAI            `True,`
	IEExtensions   TargeteNBIDExtIEs `False,OPTIONAL`
}
