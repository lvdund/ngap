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

func Test_PDUSessionResourceSetupResponseTransfer(t *testing.T) {
	msg := ies.PDUSessionResourceSetupResponseTransfer{
		DLQosFlowPerTNLInformation: &ies.QosFlowPerTNLInformation{
			UPTransportLayerInformation: &ies.UPTransportLayerInformation{
				Choice: 1,
				GTPTunnel: &ies.GTPTunnel{
					GTPTEID: &ies.GTPTEID{Value: aper.OctetString{0, 0, 0, 1}},
					TransportLayerAddress: &ies.TransportLayerAddress{Value: aper.BitString{
						Bytes:   []byte{192, 168, 57, 1},
						NumBits: 32,
					}},
				},
			},
			AssociatedQosFlowList: &ies.AssociatedQosFlowList{Value: []*ies.AssociatedQosFlowItem{&ies.AssociatedQosFlowItem{
				QosFlowIdentifier: &ies.QosFlowIdentifier{3},
			}}},
		},
	}
	var b []byte
	var err error
	if b, err = msg.Encode(); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("encode: %0b\n\t%v\n", b, b)
	}
	// need [0 3 224 192 168 57 1 0 0 0 1 0 3]
	// have [0 3 224 192 168 57 1 0 0 0 1 0 3]
	// have [0 3 224 192 168 57 1 0 0 0 1 0 3]
	// have [0 3 224 192 168 57 1 0 0 0 1 0 3]

	fmt.Println("=====================================================")
	fmt.Println("=====================================================")
	fmt.Println("=====================================================")
	fmt.Println()
	a := ies.PDUSessionResourceSetupResponseTransfer{}
	if err := a.Decode(b); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a.DLQosFlowPerTNLInformation.UPTransportLayerInformation.Choice)
		fmt.Println(a.DLQosFlowPerTNLInformation.UPTransportLayerInformation.GTPTunnel.GTPTEID)
		fmt.Println(a.DLQosFlowPerTNLInformation.UPTransportLayerInformation.GTPTunnel.TransportLayerAddress)
		fmt.Println(a.DLQosFlowPerTNLInformation.AssociatedQosFlowList.Value[0].QosFlowIdentifier)
	}

}

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

	var buf []byte
	var err error
	if buf, err = a.Encode(); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("encode: %0b\n\t%v\n", buf, buf)
	}
	fmt.Println("=====================================================")
	fmt.Println("=====================================================")
	fmt.Println("=====================================================")
	fmt.Println()
	a = ies.PDUSessionResourceSetupRequestTransfer{}
	if err, _ := a.Decode(buf); err != nil {
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

/////

func TestInitialContextSetupResponse(t *testing.T) {
	msg := ies.InitialContextSetupResponse{
		AMFUENGAPID: &ies.AMFUENGAPID{Value: 1},
		// RANUENGAPID: &ies.RANUENGAPID{Value: 1},
	}
	var buf bytes.Buffer
	if err := msg.Encode(&buf); err != nil {
		fmt.Println("err:", err)
	} else {
		b := buf.Bytes()
		fmt.Printf("encode: %0b\n\t%v\n", b, b)
	}
}

func TestPduSessionResourceSetupResponse(t *testing.T) {
	msg := ies.PDUSessionResourceSetupResponse{
		AMFUENGAPID: &ies.AMFUENGAPID{1},
		RANUENGAPID: &ies.RANUENGAPID{1},
		PDUSessionResourceSetupListSURes: &ies.PDUSessionResourceSetupListSURes{
			Value: []*ies.PDUSessionResourceSetupItemSURes{&ies.PDUSessionResourceSetupItemSURes{
				PDUSessionID:                            &ies.PDUSessionID{1},
				PDUSessionResourceSetupResponseTransfer: &aper.OctetString{0, 3, 224, 192, 168, 57, 1, 0, 0, 0, 1, 0, 3},
			}},
		},
	}
	var buf bytes.Buffer
	if err := msg.Encode(&buf); err != nil {
		fmt.Println("err:", err)
	} else {
		b := buf.Bytes()
		fmt.Printf("encode: %0b\n\t%v\n", b, b)
	}
	// [32 29 0 36 0 0 3 0 10 64 2 0 1 0 85 64 2 0 1 0 75 64 17 0 0 1 13 0 3 224 192 168 57 1 0 0 0 1 0 3]
	// [16 29 0 36 0 0 3 0 10 64 2 0 1 0 85 64 2 0 1 0 75 64 17 0 0 1 13 0 3 224 192 168 57 1 0 0 0 1 0 3]
	// [32 29 0 36 0 0 3 0 10 64 2 0 1 0 85 64 2 0 1 0 75 64 17 0 0 1 13 0 3 224 192 168 57 1 0 0 0 1 0 3]
}
