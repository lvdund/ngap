package ie

type NGResetIEs struct {
	Cause     Cause     `,ignore,mandatory`
	ResetType ResetType `,reject,mandatory`
}
