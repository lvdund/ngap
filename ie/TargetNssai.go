package ie

type TargetNssai struct {
	TargetSNssaiList *[]TargetSNssaiItem
}

type TargetSNssaiItem struct {
	SNssai *SNssai
}
