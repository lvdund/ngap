package ie

type GlobalTwifId struct {
PlmnIdentity	*PlmnIdentity
ChoiceTwifId	*ChoiceTwifId
}

type ChoiceTwifId struct {
TwifId	*TwifId
}

type TwifId struct {
TwifId	*[]byte
}
