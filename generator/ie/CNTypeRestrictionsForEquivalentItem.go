package ie

type CNTypeRestrictionsForEquivalentItem struct {
	PlmnIdentity PLMNIdentity                              `False,`
	IEExtensions CNTypeRestrictionsForEquivalentItemExtIEs `False,OPTIONAL`
}
