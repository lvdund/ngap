package ngap

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"testing"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
	"github.com/lvdund/ngap/utils"
)

func Test_transfer(t *testing.T) {
	var buff bytes.Buffer
	w := aper.NewWriter(&buff)

	msg := ies.HandoverRequiredTransfer{}
	if err := msg.Encode(w); err != nil {
		fmt.Println("err:", err)
	} else {
		w.Close()
		b := buff.Bytes()
		fmt.Printf("encode: %0b\n\t%v\n", b, b)
	}
}

func Test_PDUSessionResourceSetupRequestTransfer(t *testing.T) {
	teidOct := make([]byte, 4)
	// TODO might need to generate teid when adding real sessions
	teid := uint32(1)
	binary.BigEndian.PutUint32(teidOct, teid)
	a := ies.PDUSessionResourceSetupRequestTransfer{
		// PDUSessionAggregateMaximumBitRate: &ies.PDUSessionAggregateMaximumBitRate{
		// 	PDUSessionAggregateMaximumBitRateDL: &ies.BitRate{Value: 1},
		// 	PDUSessionAggregateMaximumBitRateUL: &ies.BitRate{Value: 1},
		// },
		ULNGUUPTNLInformation: &ies.UPTransportLayerInformation{
			Choice: ies.UPTransportLayerInformationPresentGTPTunnel,
			GTPTunnel: &ies.GTPTunnel{
				TransportLayerAddress: &ies.TransportLayerAddress{Value: aper.BitString{
					Bytes:   net.ParseIP("127.0.0.1").To4(),
					NumBits: uint64(len(net.ParseIP("127.0.0.1").To4()) * 8),
				}},
				GTPTEID: &ies.GTPTEID{Value: teidOct},
			},
		},
		PDUSessionType: &ies.PDUSessionType{Value: ies.PDUSessionTypeIpv4},
		QosFlowSetupRequestList: &ies.QosFlowSetupRequestList{Value: []*ies.QosFlowSetupRequestItem{{
			QosFlowIdentifier: &ies.QosFlowIdentifier{Value: 1},
			QosFlowLevelQosParameters: &ies.QosFlowLevelQosParameters{
				QosCharacteristics: &ies.QosCharacteristics{
					Choice: ies.QosCharacteristicsPresentNonDynamic5QI,
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

	if b, err := NgapEncode(&a); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("encode: %0b\n\t%v\n", b, b)
	}
}

func TestTransfer(t *testing.T) {
	dowlinkTeid := make([]byte, 4)
	teid := 1
	ipv4 := "192.168.1.101"
	addr := utils.IPAddressToNgap(ipv4, "")
	binary.BigEndian.PutUint32(dowlinkTeid, uint32(teid))
	fmt.Println("addr:", addr)
	fmt.Printf("dowlinkTeid: %.8b\n", dowlinkTeid)

	Transfer := ies.PDUSessionResourceSetupResponseTransfer{
		DLQosFlowPerTNLInformation: &ies.QosFlowPerTNLInformation{
			UPTransportLayerInformation: &ies.UPTransportLayerInformation{
				Choice: ies.UPTransportLayerInformationPresentGTPTunnel,
				GTPTunnel: &ies.GTPTunnel{
					GTPTEID:               &ies.GTPTEID{Value: dowlinkTeid},
					TransportLayerAddress: &addr,
				},
			},
			AssociatedQosFlowList: &ies.AssociatedQosFlowList{
				Value: []*ies.AssociatedQosFlowItem{&ies.AssociatedQosFlowItem{
					QosFlowIdentifier: &ies.QosFlowIdentifier{Value: 1},
				}},
			},
		},
	}
	_ = addr
	var buf bytes.Buffer
	r := aper.NewWriter(&buf)
	if err := Transfer.Encode(r); err != nil {
		fmt.Println("err")
	}
	res := buf.Bytes()
	fmt.Printf("encode: %0b\n%v\n", res, res)

	// buf = bytes.Buffer{}
	// r = aper.NewWriter(&buf)
	// r.WriteBool(aper.Zero)
	// op := []byte{0}
	// aper.SetBit(op, 0)
	// fmt.Println("optional:", op)
	// r.WriteBits(op, 1)
	// r.WriteBitString(addr.Value.Bytes, uint(addr.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, true)
	// fmt.Printf("========check: \t%0b\n\t\t%v\n", buf.Bytes(), buf)
	// r.WriteOctetString(dowlinkTeid, &aper.Constraint{Lb: 4, Ub: 4}, false)
	// fmt.Printf("gtp normal: \t%0b\n\n", buf.Bytes())

	// buf = bytes.Buffer{}
	// r = aper.NewWriter(&buf)
	// gtp := ies.GTPTunnel{
	// 	GTPTEID:               &ies.GTPTEID{Value: dowlinkTeid},
	// 	TransportLayerAddress: &addr,
	// }
	// if err := gtp.Encode(r); err != nil {
	// 	fmt.Println("err")
	// }
	// fmt.Printf("gtp: \t\t%0b", buf.Bytes())
}
