package ie

type LteV2XServicesAuthorized struct {
	VehicleUe    []byte `bitstring:"sizeLB:0,sizeUB:150"`
	PedestrianUe []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
