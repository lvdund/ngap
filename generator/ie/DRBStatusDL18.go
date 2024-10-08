package ie

type DRBStatusDL18 struct {
	DLCOUNTValue COUNTValueForPDCPSN18 `True,`
	IEExtension  DRBStatusDL18ExtIEs   `False,OPTIONAL`
}
