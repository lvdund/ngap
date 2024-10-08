package ie

type DRBStatusDL12 struct {
	DLCOUNTValue COUNTValueForPDCPSN12 `True,`
	IEExtension  DRBStatusDL12ExtIEs   `False,OPTIONAL`
}
