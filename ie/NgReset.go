package ie

type NgReset struct {
MessageType	*MessageType
Cause	*Cause
ChoiceResetType	*ChoiceResetType
}

type ChoiceResetType struct {
NgInterface	*NgInterface
PartOfNgInterface	*PartOfNgInterface
}

type NgInterface struct {
ResetAll	*[]byte
}

type PartOfNgInterface struct {
UeAssociatedLogicalNgConnectionList	*UeAssociatedLogicalNgConnectionList
}
