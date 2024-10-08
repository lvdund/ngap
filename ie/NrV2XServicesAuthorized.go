package ie

type NrV2XServicesAuthorized struct {
	VehicleUe    []byte `bitstring:"sizeLB:0,sizeUB:150"`
	PedestrianUe []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
