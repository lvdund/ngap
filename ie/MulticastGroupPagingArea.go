package ie

type MulticastGroupPagingArea struct {
	MbsAreaTaiList MbsAreaTaiList `bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsAreaTaiList struct {
	Tai Tai `bitstring:"sizeLB:0,sizeUB:150"`
}
