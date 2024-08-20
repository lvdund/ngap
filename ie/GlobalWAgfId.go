package ie

type GlobalWAgfId struct {
PlmnIdentity	*PlmnIdentity
ChoiceWAgfId	*ChoiceWAgfId
}

type ChoiceWAgfId struct {
WAgfId	*WAgfId
}

type WAgfId struct {
WAgfId	*[]byte
}
