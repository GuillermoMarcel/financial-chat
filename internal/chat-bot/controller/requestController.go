package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GuillermoMarcel/financial-chat/internal/shared/queue"
)

const stockUrl = "https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv"

type RequestController struct {
	CmdConsumer    queue.Consumer
	ReturnProducer queue.Producer
}

func (r RequestController) ServeApp() {
	queueChan := make(chan []byte)
	r.CmdConsumer.ReturnChan = queueChan

	r.CmdConsumer.Start()

	for msg := range queueChan {
		var request queue.StockPriceRequest

		json.Unmarshal(msg, &request)

		res := r.ExecuteQuery(request)
		if res == nil {
			internalErrorMsg := "there was a problem retrieving the stock price"
			res = &queue.StockPriceResult{
				ChatroomId: request.ChatroomId,
				UserId:     request.UserId,
				StockCode:  request.StockCode,
				Error:      &internalErrorMsg,
			}
		}

		r.ReturnProducer.SendJson(res)
	}
}

func (r RequestController) ExecuteQuery(req queue.StockPriceRequest) *queue.StockPriceResult {

	url := fmt.Sprintf(stockUrl, req.StockCode)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("url: %s\n", err.Error())
		return nil
	}
	defer resp.Body.Close()

	csvReader := csv.NewReader(resp.Body)

	_, err = csvReader.Read()
	if err != nil {
		fmt.Printf("bad headers: %s\n", err.Error())
		return nil
	}
	records, err := csvReader.Read()
	if err != nil {
		fmt.Printf("bad format: %s\n", err.Error())
		return nil
	}

	if records[6] == "N/D" {
		msg := "stock not found"
		return &queue.StockPriceResult{
			ChatroomId: req.ChatroomId,
			UserId:     req.UserId,
			StockCode:  req.StockCode,
			Error:      &msg,
		}
	}
	price, err := strconv.ParseFloat(records[6], 64)

	return &queue.StockPriceResult{
		ChatroomId: req.ChatroomId,
		UserId:     req.UserId,
		StockCode:  records[0],
		StockPrice: price,
	}
}
