package request

type Transfer struct {
	From   int64 `form:"from" json:"from,omitempty"`
	To     int64 `form:"to" json:"to,omitempty"`
	Amount int64 `form:"amount" binding:"required,gt=0" json:"amount,omitempty"`
}
