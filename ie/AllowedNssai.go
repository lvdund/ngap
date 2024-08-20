package ie

type AllowedNssai struct {
	AllowedSNssaiList *[]AllowedSNssaiItem
}

type AllowedSNssaiItem struct {
	SNssai *SNssai
}
