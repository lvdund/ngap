package ie

type SupportedTAItem struct {
	TAC               TAC                   `False,`
	BroadcastPLMNList BroadcastPLMNList     `False,`
	IEExtensions      SupportedTAItemExtIEs `False,OPTIONAL`
}
