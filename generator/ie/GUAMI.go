package ie

type GUAMI struct {
	PLMNIdentity PLMNIdentity `False,`
	AMFRegionID  AMFRegionID  `False,`
	AMFSetID     AMFSetID     `False,`
	AMFPointer   AMFPointer   `False,`
	IEExtensions GUAMIExtIEs  `False,OPTIONAL`
}
