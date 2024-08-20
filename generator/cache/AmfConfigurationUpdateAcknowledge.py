ies = []
ies.append({ "iei" : "MessageType", "level" : "0", "range" : "", "type" : "MessageType"})
ies.append({ "iei" : "AmfTnlAssociationSetupList", "level" : "0", "range" : "0..1", "type" : ""})
ies.append({ "iei" : "AmfTnlAssociationSetupItem", "level" : "1", "range" : "1..<maxnoofTNLAssociations>", "type" : ""})
ies.append({ "iei" : "AmfTnlAssociationAddress", "level" : "2", "range" : "", "type" : "CpTransportLayerInformation"})
ies.append({ "iei" : "AmfTnlAssociationFailedToSetupList", "level" : "0", "range" : "", "type" : "TnlAssociationList"})
ies.append({ "iei" : "CriticalityDiagnostics", "level" : "0", "range" : "", "type" : "CriticalityDiagnostics"})
