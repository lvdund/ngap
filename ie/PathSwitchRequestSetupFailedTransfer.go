package ie

type PathSwitchRequestSetupFailedTransfer struct {
	Cause Cause `bitstring:"sizeLB:0,sizeUB:150"`
}
