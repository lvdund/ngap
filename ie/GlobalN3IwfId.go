package ie

type GlobalN3IwfId struct {
PlmnIdentity	*PlmnIdentity
ChoiceN3IwfId	*ChoiceN3IwfId
}

type ChoiceN3IwfId struct {
N3IwfId	*N3IwfId
}

type N3IwfId struct {
N3IwfId	[]byte	//`bitstring:"sizeLB:16,sizeUB:16"`
}
