package ie

type GlobalRanNodeId struct {
ChoiceNgRanNode	*ChoiceNgRanNode
}

type ChoiceNgRanNode struct {
Gnb	*Gnb
NgEnb	*NgEnb
N3Iwf	*N3Iwf
Tngf	*Tngf
Twif	*Twif
WAgf	*WAgf
}

type Gnb struct {
GlobalGnbId	*GlobalGnbId
}

type NgEnb struct {
GlobalNgEnbId	*GlobalNgEnbId
}

type N3Iwf struct {
GlobalN3IwfId	*GlobalN3IwfId
}

type Tngf struct {
GlobalTngfId	*GlobalTngfId
}

type Twif struct {
GlobalTwifId	*GlobalTwifId
}

type WAgf struct {
GlobalWAgfId	*GlobalWAgfId
}
