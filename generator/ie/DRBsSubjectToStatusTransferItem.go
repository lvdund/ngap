package ie

type DRBsSubjectToStatusTransferItem struct {
	DRBID       DRBID                                 `False,`
	DRBStatusUL DRBStatusUL                           `False,`
	DRBStatusDL DRBStatusDL                           `False,`
	IEExtension DRBsSubjectToStatusTransferItemExtIEs `False,OPTIONAL`
}
