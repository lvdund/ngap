package ie

type UeRlfReportContainer struct {
ChoiceRlfType	*ChoiceRlfType
}

type ChoiceRlfType struct {
Nr	*Nr
Lte	*Lte
}

type Lte struct {
LteUeRlfReportContainer	*[]byte
}
