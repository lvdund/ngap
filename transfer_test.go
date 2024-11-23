package ngap

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"testing"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

func Test_PDUSessionResourceSetupRequestTransfer(t *testing.T) {
	teidOct := make([]byte, 4)
	// TODO might need to generate teid when adding real sessions
	teid := uint32(1)
	binary.BigEndian.PutUint32(teidOct, teid)
	a := ies.PDUSessionResourceSetupRequestTransfer{
		PDUSessionAggregateMaximumBitRate: &ies.PDUSessionAggregateMaximumBitRate{
			PDUSessionAggregateMaximumBitRateDL: &ies.BitRate{Value: 1},
			PDUSessionAggregateMaximumBitRateUL: &ies.BitRate{Value: 1},
		},
		ULNGUUPTNLInformation: &ies.UPTransportLayerInformation{
			Choice: ies.UPTransportLayerInformationPresentGTPTunnel,
			GTPTunnel: &ies.GTPTunnel{
				TransportLayerAddress: &ies.TransportLayerAddress{Value: aper.BitString{
					Bytes:   net.ParseIP("127.0.0.1").To4(),
					NumBits: uint64(len(net.ParseIP("127.0.0.1").To4()) * 8),
				}},
				GTPTEID: &ies.GTPTEID{Value: aper.OctetString{0, 0, 0, 1}},
			},
		},
		PDUSessionType: &ies.PDUSessionType{Value: ies.PDUSessionTypeIpv4},
		QosFlowSetupRequestList: &ies.QosFlowSetupRequestList{Value: []*ies.QosFlowSetupRequestItem{{
			QosFlowIdentifier: &ies.QosFlowIdentifier{Value: 1},
			QosFlowLevelQosParameters: &ies.QosFlowLevelQosParameters{
				QosCharacteristics: &ies.QosCharacteristics{
					Choice:        ies.QosCharacteristicsPresentNonDynamic5QI,
					NonDynamic5QI: &ies.NonDynamic5QIDescriptor{FiveQI: &ies.FiveQI{Value: 1}},
				},
				AllocationAndRetentionPriority: &ies.AllocationAndRetentionPriority{
					PriorityLevelARP: &ies.PriorityLevelARP{Value: 1},
					PreemptionCapability: &ies.PreemptionCapability{
						Value: ies.PreemptionCapabilityShallnottriggerpreemption,
					},
					PreemptionVulnerability: &ies.PreemptionVulnerability{
						Value: ies.PreemptionCapabilityShallnottriggerpreemption,
					},
				},
			},
		}}},
	}

	var buf bytes.Buffer
	if err := a.Encode(&buf); err != nil {
		fmt.Println("err:", err)
	} else {
		b := buf.Bytes()
		fmt.Printf("encode: %0b\n\t%v\n", b, b)
	}
	fmt.Println("=====================================================")
	fmt.Println("=====================================================")
	fmt.Println("=====================================================")
	fmt.Println()
	b := buf.Bytes()
	a = ies.PDUSessionResourceSetupRequestTransfer{}
	if err, _ := a.Decode(b); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
		fmt.Println(a.PDUSessionAggregateMaximumBitRate.PDUSessionAggregateMaximumBitRateDL)
		fmt.Println(a.PDUSessionAggregateMaximumBitRate.PDUSessionAggregateMaximumBitRateUL)
		fmt.Println(a.ULNGUUPTNLInformation.Choice)
		fmt.Println(a.ULNGUUPTNLInformation.GTPTunnel.TransportLayerAddress.Value)
		fmt.Println(a.ULNGUUPTNLInformation.GTPTunnel.GTPTEID.Value)
		fmt.Println(a.PDUSessionType)
		fmt.Println(a.QosFlowSetupRequestList.Value[0].QosFlowIdentifier.Value)
		fmt.Println(a.QosFlowSetupRequestList.Value[0].QosFlowLevelQosParameters.QosCharacteristics.Choice)
		fmt.Println(a.QosFlowSetupRequestList.Value[0].QosFlowLevelQosParameters.QosCharacteristics.NonDynamic5QI.FiveQI.Value)
		fmt.Println(a.QosFlowSetupRequestList.Value[0].QosFlowLevelQosParameters.AllocationAndRetentionPriority.PriorityLevelARP)
		fmt.Println(a.QosFlowSetupRequestList.Value[0].QosFlowLevelQosParameters.AllocationAndRetentionPriority.PreemptionCapability)
		fmt.Println(a.QosFlowSetupRequestList.Value[0].QosFlowLevelQosParameters.AllocationAndRetentionPriority.PreemptionVulnerability)

	}
}
