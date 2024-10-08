package ie

type UserPlaneSecurityInformation struct {
	SecurityResult     SecurityResult                     `True,`
	SecurityIndication SecurityIndication                 `True,`
	IEExtensions       UserPlaneSecurityInformationExtIEs `False,OPTIONAL`
}
