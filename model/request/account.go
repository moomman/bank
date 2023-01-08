package request

type Transfer struct {
	From   int64 `form:"from" json:"from,omitempty"`
	To     int64 `form:"to" json:"to,omitempty"`
	Amount int64 `form:"amount" binding:"required,gt=0" json:"amount,omitempty"`
}

type AddAccount struct {
	Owner    string `form:"owner" binding:"required" json:"owner"`
	Balance  int64  `form:"balance" binding:"required" json:"balance"`
	Currency string `form:"currency" binding:"required" json:"currency"`
}
