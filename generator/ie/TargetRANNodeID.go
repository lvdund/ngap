package ie

type TargetRANNodeID struct {
	GlobalRANNodeID GlobalRANNodeID       `False,`
	SelectedTAI     TAI                   `True,`
	IEExtensions    TargetRANNodeIDExtIEs `False,OPTIONAL`
}
