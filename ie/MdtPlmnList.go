package ie

type MdtPlmnList struct {
	PlmnIdentity PlmnIdentity `bitstring:"sizeLB:0,sizeUB:150"`
}
