package ie

const (
	InitiatingMessagePresentNothing int = iota /* No components present */
	InitiatingMessagePresentAMFConfigurationUpdate
	InitiatingMessagePresentHandoverCancel
	InitiatingMessagePresentHandoverRequired
	InitiatingMessagePresentHandoverRequest
	InitiatingMessagePresentInitialContextSetupRequest
	InitiatingMessagePresentNGReset
	InitiatingMessagePresentNGSetupRequest
	InitiatingMessagePresentPathSwitchRequest
	InitiatingMessagePresentPDUSessionResourceModifyRequest
	InitiatingMessagePresentPDUSessionResourceModifyIndication
	InitiatingMessagePresentPDUSessionResourceReleaseCommand
	InitiatingMessagePresentPDUSessionResourceSetupRequest
	InitiatingMessagePresentPWSCancelRequest
	InitiatingMessagePresentRANConfigurationUpdate
	InitiatingMessagePresentUEContextModificationRequest
	InitiatingMessagePresentUEContextReleaseCommand
	InitiatingMessagePresentUERadioCapabilityCheckRequest
	InitiatingMessagePresentWriteReplaceWarningRequest
	InitiatingMessagePresentAMFStatusIndication
	InitiatingMessagePresentCellTrafficTrace
	InitiatingMessagePresentDeactivateTrace
	InitiatingMessagePresentDownlinkNASTransport
	InitiatingMessagePresentDownlinkNonUEAssociatedNRPPaTransport
	InitiatingMessagePresentDownlinkRANConfigurationTransfer
	InitiatingMessagePresentDownlinkRANStatusTransfer
	InitiatingMessagePresentDownlinkUEAssociatedNRPPaTransport
	InitiatingMessagePresentErrorIndication
	InitiatingMessagePresentHandoverNotify
	InitiatingMessagePresentInitialUEMessage
	InitiatingMessagePresentLocationReport
	InitiatingMessagePresentLocationReportingControl
	InitiatingMessagePresentLocationReportingFailureIndication
	InitiatingMessagePresentNASNonDeliveryIndication
	InitiatingMessagePresentOverloadStart
	InitiatingMessagePresentOverloadStop
	InitiatingMessagePresentPaging
	InitiatingMessagePresentPDUSessionResourceNotify
	InitiatingMessagePresentPrivateMessage
	InitiatingMessagePresentPWSFailureIndication
	InitiatingMessagePresentPWSRestartIndication
	InitiatingMessagePresentRerouteNASRequest
	InitiatingMessagePresentRRCInactiveTransitionReport
	InitiatingMessagePresentSecondaryRATDataUsageReport
	InitiatingMessagePresentTraceFailureIndication
	InitiatingMessagePresentTraceStart
	InitiatingMessagePresentUEContextReleaseRequest
	InitiatingMessagePresentUERadioCapabilityInfoIndication
	InitiatingMessagePresentUETNLABindingReleaseRequest
	InitiatingMessagePresentUplinkNASTransport
	InitiatingMessagePresentUplinkNonUEAssociatedNRPPaTransport
	InitiatingMessagePresentUplinkRANConfigurationTransfer
	InitiatingMessagePresentUplinkRANStatusTransfer
	InitiatingMessagePresentUplinkUEAssociatedNRPPaTransport
)