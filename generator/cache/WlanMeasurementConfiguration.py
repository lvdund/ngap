ies = []
ies.append({ "iei" : "WlanMeasurementConfiguration", "level" : "0", "range" : "", "type" : "ENUMERATED (Setup, …)"})
ies.append({ "iei" : "WlanMeasurementConfigurationNameList", "level" : "0", "range" : "0..1", "type" : ""})
ies.append({ "iei" : "WlanMeasurementConfigurationNameItem", "level" : "1", "range" : "1..<maxnoofWLANName>", "type" : ""})
ies.append({ "iei" : "WlanMeasurementConfigurationName", "level" : "2", "range" : "", "type" : "OCTET STRING (SIZE (1..32))"})
ies.append({ "iei" : "WlanRssi", "level" : "0", "range" : "", "type" : "ENUMERATED (true, …)"})
ies.append({ "iei" : "WlanRtt", "level" : "0", "range" : "", "type" : "ENUMERATED (true, …)"})
