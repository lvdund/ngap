ies = []
ies.append({ "iei" : "ProcedureCode", "level" : "0", "range" : "", "type" : "INTEGER (0..255)"})
ies.append({ "iei" : "TriggeringMessage", "level" : "0", "range" : "", "type" : "ENUMERATED (initiating message, successful outcome, unsuccessful outcome)"})
ies.append({ "iei" : "ProcedureCriticality", "level" : "0", "range" : "", "type" : "ENUMERATED (reject, ignore, notify)"})
ies.append({ "iei" : "InformationElementCriticalityDiagnostics", "level" : "0", "range" : "0..<maxnoofErrors>", "type" : ""})
ies.append({ "iei" : "IeCriticality", "level" : "1", "range" : "", "type" : "ENUMERATED (reject, ignore, notify)"})
ies.append({ "iei" : "IeId", "level" : "1", "range" : "", "type" : "INTEGER (0..65535)"})
ies.append({ "iei" : "TypeOfError", "level" : "1", "range" : "", "type" : "ENUMERATED (not understood, missing, …)"})
