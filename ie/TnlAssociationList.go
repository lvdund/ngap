package ie

type TnlAssociationList struct {
TnlAssociationItem	*TnlAssociationItem
}

type TnlAssociationItem struct {
TnlAssociationAddress	*CpTransportLayerInformation
Cause	*Cause
}
