package ie

import "ngap/aper"

type ExpectedUeBehaviour struct {
	ExpectedUeActivityBehaviour ExpectedUeActivityBehaviour //`bitstring:"sizeLB:0,sizeUB:150"`
	ExpectedHoInterval          []byte                      //`bitstring:"sizeLB:0,sizeUB:150"`
	ExpectedUeMobility          []byte                      //`bitstring:"sizeLB:0,sizeUB:150"`
	ExpectedUeMovingTrajectory  ExpectedUeMovingTrajectory  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ExpectedUeMovingTrajectory struct {
	ExpectedUeMovingTrajectoryItem ExpectedUeMovingTrajectoryItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ExpectedUeMovingTrajectoryItem struct {
	NgRanCgi         NgRanCgi     //`bitstring:"sizeLB:0,sizeUB:150"`
	TimeStayedInCell aper.Integer //`Integer:"valueLB:0,valueUB:4095"`
}
