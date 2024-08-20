package ie

type NrNtnTaiInformation struct {
ServingPlmn	*PlmnIdentity
TacListInNrNtn	*TacListInNrNtn
UeLocationDerivedTacInNrNtn	*Tac
}

type TacListInNrNtn struct {
Tac	*Tac
}
