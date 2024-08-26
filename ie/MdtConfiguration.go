package ie

type MdtConfiguration struct {
	MdtConfigurationNr    MdtConfigurationNr    `bitstring:"sizeLB:0,sizeUB:150"`
	MdtConfigurationEutra MdtConfigurationEutra `bitstring:"sizeLB:0,sizeUB:150"`
}
