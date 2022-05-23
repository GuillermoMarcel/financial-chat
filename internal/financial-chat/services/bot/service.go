package bot

import (
	"encoding/json"
	"log"

	"github.com/GuillermoMarcel/financial-chat/internal/shared/queue"
)

type Service struct {
	Log             *log.Logger
	CmdProducer     *queue.Producer
	ResultsConsumer *queue.Consumer
	ResultChan      chan queue.StockPriceResult
}

func (s Service) RequestStockPrice(chatroomId uint, userId uint, code string) {
	req := queue.StockPriceRequest{
		ChatroomId: chatroomId,
		UserId:     userId,
		StockCode:  code,
	}

	s.CmdProducer.SendJson(req)
}

func (s Service) ReadIncoming() {
	s.ResultsConsumer.Start()
	defer s.ResultsConsumer.Stop()

	log.Println("reading consumer results")

	reader := make(chan []byte)
	s.ResultsConsumer.ReturnChan = reader

	for msg := range reader {
		var result queue.StockPriceResult
		json.Unmarshal(msg, &result)

		s.ResultChan <- result
	}
}
