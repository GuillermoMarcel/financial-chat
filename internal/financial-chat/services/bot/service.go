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

func (s *Service) RequestStockPrice(chatroomId uint, userId uint, code string) {
	req := queue.StockPriceRequest{
		ChatroomId: chatroomId,
		UserId:     userId,
		StockCode:  code,
	}

	s.CmdProducer.SendJson(req)
}

func (s *Service) ReadIncoming() {
	reader := make(chan []byte)
	s.ResultsConsumer.ReturnChan = reader

	s.ResultsConsumer.Start()
	defer s.ResultsConsumer.Stop()

	s.Log.Println("reading consumer results")

	for msg := range reader {

		var result queue.StockPriceResult
		json.Unmarshal(msg, &result)
		if s.ResultChan == nil {
			s.Log.Printf("result channel not set, mesasge lost %v\n", result)
		}
		s.ResultChan <- result
	}
}
