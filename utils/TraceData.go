package utils

import (
	"encoding/hex"
	"strings"

	"github.com/lvdund/ngap/ies"
	"github.com/sirupsen/logrus"
)

type TraceDepth string

// List of TraceDepth
const (
	TraceDepth_MINIMUM                     TraceDepth = "MINIMUM"
	TraceDepth_MEDIUM                      TraceDepth = "MEDIUM"
	TraceDepth_MAXIMUM                     TraceDepth = "MAXIMUM"
	TraceDepth_MINIMUM_WO_VENDOR_EXTENSION TraceDepth = "MINIMUM_WO_VENDOR_EXTENSION"
	TraceDepth_MEDIUM_WO_VENDOR_EXTENSION  TraceDepth = "MEDIUM_WO_VENDOR_EXTENSION"
	TraceDepth_MAXIMUM_WO_VENDOR_EXTENSION TraceDepth = "MAXIMUM_WO_VENDOR_EXTENSION"
)

type TraceData struct {
	TraceRef                 string
	TraceDepth               TraceDepth
	NeTypeList               string
	EventList                string
	CollectionEntityIpv4Addr string
	CollectionEntityIpv6Addr string
	InterfaceList            string
}

func TraceDataToModels(traceActivation ies.TraceActivation) (traceData TraceData) {
	// TODO: finish this function when need
	return
}

func TraceDataToNgap(traceData TraceData, trsr string) ies.TraceActivation {
	var traceActivation ies.TraceActivation

	if len(trsr) != 4 {
		logrus.Warningln("Trace Recording Session Reference should be 2 octets")
		return traceActivation
	}

	// NG-RAN Trace ID (left most 6 octet Trace Reference + last 2 octet Trace Recoding Session Reference)
	subStringSlice := strings.Split(traceData.TraceRef, "-")

	if len(subStringSlice) != 2 {
		logrus.Warningln("TraceRef format is not correct")
		return traceActivation
	}

	plmnID := PlmnId{}
	plmnID.Mcc = subStringSlice[0][:3]
	plmnID.Mnc = subStringSlice[0][3:]
	var traceID []byte
	if traceIDTmp, err := hex.DecodeString(subStringSlice[1]); err != nil {
		logrus.Warnf("")
	} else {
		traceID = traceIDTmp
	}

	tmp := PlmnIdToNgap(plmnID)
	traceReference := append(tmp, traceID...)
	var trsrNgap []byte
	if trsrNgapTmp, err := hex.DecodeString(trsr); err != nil {
		logrus.Warnf("Decode trsr failed: %+v", err)
	} else {
		trsrNgap = trsrNgapTmp
	}

	nGRANTraceID := append(traceReference, trsrNgap...)

	traceActivation.NGRANTraceID = nGRANTraceID

	// Interfaces To Trace
	var interfacesToTrace []byte
	if interfacesToTraceTmp, err := hex.DecodeString(traceData.InterfaceList); err != nil {
		logrus.Warnf("Decode Interface failed: %+v", err)
	} else {
		interfacesToTrace = interfacesToTraceTmp
	}
	traceActivation.InterfacesToTrace = interfacesToTrace

	// Trace Collection Entity IP Address
	ngapIP := IPAddressToNgap(traceData.CollectionEntityIpv4Addr, traceData.CollectionEntityIpv6Addr)
	traceActivation.TraceCollectionEntityIPAddress = ngapIP

	// Trace Depth
	switch traceData.TraceDepth {
	case TraceDepth_MINIMUM:
		traceActivation.TraceDepth.Value = ies.TraceDepthMinimum
	case TraceDepth_MEDIUM:
		traceActivation.TraceDepth.Value = ies.TraceDepthMedium
	case TraceDepth_MAXIMUM:
		traceActivation.TraceDepth.Value = ies.TraceDepthMaximum
	case TraceDepth_MINIMUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ies.TraceDepthMinimumwithoutvendorspecificextension
	case TraceDepth_MEDIUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ies.TraceDepthMediumwithoutvendorspecificextension
	case TraceDepth_MAXIMUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ies.TraceDepthMaximumwithoutvendorspecificextension
	}

	return traceActivation
}
