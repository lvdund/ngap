ies = []
ies.append({ "iei" : "EventType", "level" : "0", "range" : "", "type" : "ENUMERATED (direct, change of serving cell, UE presence in the area of interest, stop change of serving cell, stop UE presence in the area of interest, cancel location reporting for the UE, …)"})
ies.append({ "iei" : "ReportArea", "level" : "0", "range" : "", "type" : "ENUMERATED (cell, …)"})
ies.append({ "iei" : "AreaOfInterestList", "level" : "0", "range" : "0..1", "type" : ""})
ies.append({ "iei" : "AreaOfInterestItem", "level" : "1", "range" : "1..<maxnoofAoI>", "type" : ""})
ies.append({ "iei" : "AreaOfInterest", "level" : "2", "range" : "", "type" : "AreaOfInterest"})
ies.append({ "iei" : "LocationReportingReferenceId", "level" : "2", "range" : "", "type" : "LocationReportingReferenceId"})
ies.append({ "iei" : "LocationReportingReferenceIdToBeCancelled", "level" : "0", "range" : "", "type" : "LocationReportingReferenceId"})
ies.append({ "iei" : "AdditionalLocationInformation", "level" : "0", "range" : "", "type" : "ENUMERATED (Include PSCell, ...)"})
