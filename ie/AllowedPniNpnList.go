package ie

type AllowedPniNpnList struct {
	AllowedPniNpnItem *AllowedPniNpnItem
}

type AllowedPniNpnItem struct {
	PlmnIdentity          *PlmnIdentity
	PniNpnRestricted      *[]byte
	AllowedCagListPerPlmn *AllowedCagListPerPlmn
}

type AllowedCagListPerPlmn struct {
	CagId *CagId
}
