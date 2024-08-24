package ie

type NrNtnTaiInformation struct {
	ServingPlmn                 PlmnIdentity   //`bitstring:"sizeLB:0,sizeUB:150"`
	TacListInNrNtn              TacListInNrNtn //`bitstring:"sizeLB:0,sizeUB:150"`
	UeLocationDerivedTacInNrNtn Tac            //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TacListInNrNtn struct {
	Tac Tac //`bitstring:"sizeLB:0,sizeUB:150"`
}
