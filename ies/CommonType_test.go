package ies

import (
	"bytes"
	"testing"

	"github.com/lvdund/ngap/aper"
)

func encoded(t *testing.T, fn func(*aper.AperWriter) error) []byte {
	t.Helper()
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err := fn(w); err != nil {
		t.Fatalf("encode: %v", err)
	}
	w.Close()
	return buf.Bytes()
}

func TestTacRoundTrip(t *testing.T) {
	item := &ServiceAreaInformationItem{
		PLMNIdentity: []byte{0x02, 0xf8, 0x39},
		AllowedTACs:  []TAC{{Value: []byte{0xaa, 0xbb, 0xcc}}},
	}
	got := encoded(t, item.Encode)
	want := []byte{0x40, 0x02, 0xf8, 0x39, 0x00, 0xaa, 0xbb, 0xcc}
	if !bytes.Equal(got, want) {
		t.Fatalf("encoded % x, want % x", got, want)
	}

	back := new(ServiceAreaInformationItem)
	if err := back.Decode(aper.NewReader(bytes.NewReader(got))); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(back.AllowedTACs) != 1 || !bytes.Equal(back.AllowedTACs[0].Value, []byte{0xaa, 0xbb, 0xcc}) {
		t.Fatalf("decoded AllowedTACs %+v", back.AllowedTACs)
	}
	if len(back.NotAllowedTACs) != 0 {
		t.Fatalf("decoded NotAllowedTACs %+v", back.NotAllowedTACs)
	}
}

func TestTacWrongLengthIsRefused(t *testing.T) {
	item := &ForbiddenAreaInformationItem{
		PLMNIdentity:  []byte{0x02, 0xf8, 0x39},
		ForbiddenTACs: []TAC{{Value: []byte{0xaa, 0xbb}}},
	}
	var buf bytes.Buffer
	if err := item.Encode(aper.NewWriter(&buf)); err == nil {
		t.Fatalf("a two-octet TAC must not encode")
	}
}

func TestMobilityRestrictionListRoundTrip(t *testing.T) {
	mrl := &MobilityRestrictionList{
		ServingPLMN:     []byte{0x02, 0xf8, 0x39},
		EquivalentPLMNs: []PLMNIdentity{{Value: []byte{0x02, 0xf8, 0x99}}},
		ForbiddenAreaInformation: []ForbiddenAreaInformationItem{
			{PLMNIdentity: []byte{0x02, 0xf8, 0x39}, ForbiddenTACs: []TAC{{Value: []byte{0x00, 0x00, 0x09}}}},
		},
		ServiceAreaInformation: []ServiceAreaInformationItem{
			{PLMNIdentity: []byte{0x02, 0xf8, 0x39}, NotAllowedTACs: []TAC{{Value: []byte{0x00, 0x00, 0x02}}}},
		},
	}
	got := encoded(t, mrl.Encode)
	back := new(MobilityRestrictionList)
	if err := back.Decode(aper.NewReader(bytes.NewReader(got))); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(back.EquivalentPLMNs) != 1 || !bytes.Equal(back.EquivalentPLMNs[0].Value, []byte{0x02, 0xf8, 0x99}) {
		t.Fatalf("decoded EquivalentPLMNs %+v", back.EquivalentPLMNs)
	}
	if len(back.ForbiddenAreaInformation) != 1 ||
		!bytes.Equal(back.ForbiddenAreaInformation[0].ForbiddenTACs[0].Value, []byte{0x00, 0x00, 0x09}) {
		t.Fatalf("decoded ForbiddenAreaInformation %+v", back.ForbiddenAreaInformation)
	}
	if len(back.ServiceAreaInformation) != 1 ||
		!bytes.Equal(back.ServiceAreaInformation[0].NotAllowedTACs[0].Value, []byte{0x00, 0x00, 0x02}) {
		t.Fatalf("decoded ServiceAreaInformation %+v", back.ServiceAreaInformation)
	}
}

// the shape the type exists for: a SEQUENCE OF whose elements are these
func TestEmergencyAreaIDSequence(t *testing.T) {
	ids := []*EmergencyAreaID{
		{Value: []byte{0x0a, 0x0b, 0x0c}},
		{Value: []byte{0x0d, 0x0e, 0x0f}},
	}
	seq := NewSequence[*EmergencyAreaID](ids, aper.Constraint{Lb: 1, Ub: maxnoofEAIforRestart}, false)
	got := encoded(t, seq.Encode)
	for _, id := range ids {
		if !bytes.Contains(got, id.Value) {
			t.Fatalf("emergency area id % x missing from % x", id.Value, got)
		}
	}

	back := NewSequence[*EmergencyAreaID](nil, aper.Constraint{Lb: 1, Ub: maxnoofEAIforRestart}, false)
	fn := func() *EmergencyAreaID { return new(EmergencyAreaID) }
	if err := back.Decode(aper.NewReader(bytes.NewReader(got)), fn); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(back.Value) != len(ids) {
		t.Fatalf("decoded %d ids, want %d (wire % x)", len(back.Value), len(ids), got)
	}
	for i, want := range ids {
		if !bytes.Equal(back.Value[i].Value, want.Value) {
			t.Fatalf("id %d decoded % x, want % x", i, back.Value[i].Value, want.Value)
		}
	}
}

func TestEmergencyAreaIDElement(t *testing.T) {
	id := &EmergencyAreaID{Value: []byte{0x0a, 0x0b, 0x0c}}
	got := encoded(t, id.Encode)
	if !bytes.Equal(got, []byte{0x0a, 0x0b, 0x0c}) {
		t.Fatalf("encoded % x, want 0a 0b 0c", got)
	}
	back := new(EmergencyAreaID)
	if err := back.Decode(aper.NewReader(bytes.NewReader(got))); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !bytes.Equal(back.Value, id.Value) {
		t.Fatalf("decoded % x, want % x", back.Value, id.Value)
	}
	short := &EmergencyAreaID{Value: []byte{0x0a, 0x0b}}
	var buf bytes.Buffer
	if err := short.Encode(aper.NewWriter(&buf)); err == nil {
		t.Fatalf("a two-octet emergency area id must not encode")
	}
}

func TestTransportLayerAddressRoundTrip(t *testing.T) {
	//an IPv4 address, an IPv6 address and both together: the three lengths the
	//protocol actually carries
	for _, addr := range [][]byte{
		{10, 0, 0, 1},
		bytes.Repeat([]byte{0xfe}, 16),
		bytes.Repeat([]byte{0xab}, 20),
	} {
		tla := &TransportLayerAddress{Value: addr}
		got := encoded(t, tla.Encode)
		back := new(TransportLayerAddress)
		if err := back.Decode(aper.NewReader(bytes.NewReader(got))); err != nil {
			t.Fatalf("%d octets: decode: %v", len(addr), err)
		}
		if !bytes.Equal(back.Value, addr) {
			t.Fatalf("%d octets: decoded % x, want % x (wire % x)", len(addr), back.Value, addr, got)
		}
	}
}

// the same value encoded as the bit string the neighbouring field in
// XnExtTLAItem declares must produce identical bytes
func TestTransportLayerAddressMatchesIPsecTLA(t *testing.T) {
	addr := []byte{10, 0, 0, 1}
	tla := &TransportLayerAddress{Value: addr}
	asBitString := NewBITSTRING(aper.BitString{Bytes: addr, NumBits: 32},
		aper.Constraint{Lb: 1, Ub: 160}, true)
	if a, b := encoded(t, tla.Encode), encoded(t, asBitString.Encode); !bytes.Equal(a, b) {
		t.Fatalf("TransportLayerAddress % x, IPsecTLA-style bit string % x", a, b)
	}
}

func TestXnTNLConfigurationInfoRoundTrip(t *testing.T) {
	info := &XnTNLConfigurationInfo{
		XnTransportLayerAddresses: []TransportLayerAddress{
			{Value: []byte{10, 0, 0, 1}},
			{Value: []byte{10, 0, 0, 2}},
		},
	}
	got := encoded(t, info.Encode)
	back := new(XnTNLConfigurationInfo)
	if err := back.Decode(aper.NewReader(bytes.NewReader(got))); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(back.XnTransportLayerAddresses) != 2 {
		t.Fatalf("decoded %d addresses (wire % x)", len(back.XnTransportLayerAddresses), got)
	}
	for i, want := range info.XnTransportLayerAddresses {
		if !bytes.Equal(back.XnTransportLayerAddresses[i].Value, want.Value) {
			t.Fatalf("address %d decoded % x, want % x", i, back.XnTransportLayerAddresses[i].Value, want.Value)
		}
	}
}
