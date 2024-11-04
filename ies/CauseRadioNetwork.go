package ies

import "github.com/lvdund/ngap/aper"

const (
	CauseRadioNetworkUnspecified                                              aper.Enumerated = 0
	CauseRadioNetworkTxnrelocoverallexpiry                                    aper.Enumerated = 1
	CauseRadioNetworkSuccessfulhandover                                       aper.Enumerated = 2
	CauseRadioNetworkReleaseduetongrangeneratedreason                         aper.Enumerated = 3
	CauseRadioNetworkReleasedueto5Gcgeneratedreason                           aper.Enumerated = 4
	CauseRadioNetworkHandovercancelled                                        aper.Enumerated = 5
	CauseRadioNetworkPartialhandover                                          aper.Enumerated = 6
	CauseRadioNetworkHofailureintarget5Gcngrannodeortargetsystem              aper.Enumerated = 7
	CauseRadioNetworkHotargetnotallowed                                       aper.Enumerated = 8
	CauseRadioNetworkTngrelocoverallexpiry                                    aper.Enumerated = 9
	CauseRadioNetworkTngrelocprepexpiry                                       aper.Enumerated = 10
	CauseRadioNetworkCellnotavailable                                         aper.Enumerated = 11
	CauseRadioNetworkUnknowntargetid                                          aper.Enumerated = 12
	CauseRadioNetworkNoradioresourcesavailableintargetcell                    aper.Enumerated = 13
	CauseRadioNetworkUnknownlocaluengapid                                     aper.Enumerated = 14
	CauseRadioNetworkInconsistentremoteuengapid                               aper.Enumerated = 15
	CauseRadioNetworkHandoverdesirableforradioreason                          aper.Enumerated = 16
	CauseRadioNetworkTimecriticalhandover                                     aper.Enumerated = 17
	CauseRadioNetworkResourceoptimisationhandover                             aper.Enumerated = 18
	CauseRadioNetworkReduceloadinservingcell                                  aper.Enumerated = 19
	CauseRadioNetworkUserinactivity                                           aper.Enumerated = 20
	CauseRadioNetworkRadioconnectionwithuelost                                aper.Enumerated = 21
	CauseRadioNetworkRadioresourcesnotavailable                               aper.Enumerated = 22
	CauseRadioNetworkInvalidqoscombination                                    aper.Enumerated = 23
	CauseRadioNetworkFailureinradiointerfaceprocedure                         aper.Enumerated = 24
	CauseRadioNetworkInteractionwithotherprocedure                            aper.Enumerated = 25
	CauseRadioNetworkUnknownpdusessionid                                      aper.Enumerated = 26
	CauseRadioNetworkUnkownqosflowid                                          aper.Enumerated = 27
	CauseRadioNetworkMultiplepdusessionidinstances                            aper.Enumerated = 28
	CauseRadioNetworkMultipleqosflowidinstances                               aper.Enumerated = 29
	CauseRadioNetworkEncryptionandorintegrityprotectionalgorithmsnotsupported aper.Enumerated = 30
	CauseRadioNetworkNgintrasystemhandovertriggered                           aper.Enumerated = 31
	CauseRadioNetworkNgintersystemhandovertriggered                           aper.Enumerated = 32
	CauseRadioNetworkXnhandovertriggered                                      aper.Enumerated = 33
	CauseRadioNetworkNotsupported5Qivalue                                     aper.Enumerated = 34
	CauseRadioNetworkUecontexttransfer                                        aper.Enumerated = 35
	CauseRadioNetworkImsvoiceepsfallbackorratfallbacktriggered                aper.Enumerated = 36
	CauseRadioNetworkUpintegrityprotectionnotpossible                         aper.Enumerated = 37
	CauseRadioNetworkUpconfidentialityprotectionnotpossible                   aper.Enumerated = 38
	CauseRadioNetworkSlicenotsupported                                        aper.Enumerated = 39
	CauseRadioNetworkUeinrrcinactivestatenotreachable                         aper.Enumerated = 40
	CauseRadioNetworkRedirection                                              aper.Enumerated = 41
	CauseRadioNetworkResourcesnotavailablefortheslice                         aper.Enumerated = 42
	CauseRadioNetworkUemaxintegrityprotecteddataratereason                    aper.Enumerated = 43
	CauseRadioNetworkReleaseduetocndetectedmobility                           aper.Enumerated = 44
	CauseRadioNetworkN26Interfacenotavailable                                 aper.Enumerated = 45
	CauseRadioNetworkReleaseduetopreemption                                   aper.Enumerated = 46
)

type CauseRadioNetwork struct {
	Value aper.Enumerated `True,0,46`
}

func (ie *CauseRadioNetwork) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 46}, true)
	return
}
func (ie *CauseRadioNetwork) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 46}, true)
	ie.Value = aper.Enumerated(v)
	return
}
