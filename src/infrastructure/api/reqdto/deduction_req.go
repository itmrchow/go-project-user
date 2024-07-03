package reqdto

type DeductionReq struct {
	walletId    uint
	amount      float64
	eventName   string
	description string
}
