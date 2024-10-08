package ie

type SourceRANNodeID struct {
	GlobalRANNodeID GlobalRANNodeID       `False,`
	SelectedTAI     TAI                   `True,`
	IEExtensions    SourceRANNodeIDExtIEs `False,OPTIONAL`
}
