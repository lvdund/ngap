package ie

type CriticalityDiagnostics struct {
	ProcedureCode             ProcedureCode                `False,OPTIONAL`
	TriggeringMessage         TriggeringMessage            `False,OPTIONAL`
	ProcedureCriticality      Criticality                  `False,OPTIONAL`
	IEsCriticalityDiagnostics CriticalityDiagnosticsIEList `False,OPTIONAL`
	IEExtensions              CriticalityDiagnosticsExtIEs `False,OPTIONAL`
}
