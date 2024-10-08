package ie

type QosFlowNotifyItem struct {
	QosFlowIdentifier QosFlowIdentifier       `False,`
	NotificationCause NotificationCause       `False,`
	IEExtensions      QosFlowNotifyItemExtIEs `False,OPTIONAL`
}
