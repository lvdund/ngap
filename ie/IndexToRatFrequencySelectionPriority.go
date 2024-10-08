package ie

import "ngap/aper"

type IndexToRatFrequencySelectionPriority struct {
	IndexToRatFrequencySelectionPriority aper.Integer `Integer:"valueLB:1,valueUB:256"`
}
