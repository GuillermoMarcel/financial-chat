package queue

type StockPriceRequest struct {
	ChatroomId uint
	UserId     uint
	StockCode  string
}

type StockPriceResult struct {
	ChatroomId uint
	UserId     uint
	StockCode  string
	StockPrice float64
	Error      *string
}
