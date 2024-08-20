package ie

type InterSystemCellActivationReply struct {
	ActivatedCellsList *ActivatedCellsList
	ActivationId       uint16 //`bitstring:"sizeLB:0,sizeUB:16384"`
}

type ActivatedCellsList struct {
	NgRanCgi *NgRanCgi
}
