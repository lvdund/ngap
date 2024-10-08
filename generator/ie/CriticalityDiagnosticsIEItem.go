package ie

type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality                        `False,`
	IEID          ProtocolIEID                       `False,`
	TypeOfError   TypeOfError                        `False,`
	IEExtensions  CriticalityDiagnosticsIEItemExtIEs `False,OPTIONAL`
}
