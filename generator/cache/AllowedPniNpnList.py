ies = []
ies.append({ "iei" : "AllowedPniNpnItem", "level" : "0", "range" : "1..<maxnoofEPLMNs+1>", "type" : ""})
ies.append({ "iei" : "PlmnIdentity", "level" : "1", "range" : "", "type" : "PlmnIdentity"})
ies.append({ "iei" : "PniNpnRestricted", "level" : "1", "range" : "", "type" : "ENUMERATED (restricted, not-restricted, …)"})
ies.append({ "iei" : "AllowedCagListPerPlmn", "level" : "1", "range" : "1..<maxnoofAllowedCAGsperPLMN>", "type" : ""})
ies.append({ "iei" : "CagId", "level" : "2", "range" : "", "type" : "CagId"})
