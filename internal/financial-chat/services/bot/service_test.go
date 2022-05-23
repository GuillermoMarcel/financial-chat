package bot

import (
	"log"
	"testing"
	"time"

	"github.com/GuillermoMarcel/financial-chat/internal/shared/queue"
)

type MockConsumer struct {
	RegisteredChan chan []byte
}

func (m MockConsumer) Start()                             {}
func (m MockConsumer) Stop()                              {}
func (m *MockConsumer) RegisterReturnChan(ch chan []byte) { m.RegisteredChan = ch }

func TestReciving(t *testing.T) {

	retCh := make(chan queue.StockPriceResult)
	defer close(retCh)

	mockCons := &MockConsumer{}
	t.Logf("%p\n", mockCons)

	service := Service{
		Log:             testLogger(t),
		ResultsConsumer: mockCons,
		ResultChan:      retCh,
	}

	go service.ReadIncoming()
	time.Sleep(time.Millisecond * 100)

	mockCons.RegisteredChan <- []byte("{\"ChatroomId\":1,\"UserId\":1,\"StockCode\":\"AAPL.US\",\"StockPrice\":143.11,\"Error\":null}")

	result := <-retCh

	if result.ChatroomId != 1 || result.StockPrice != 143.11 {
		t.Error("Failed to recive message")
	}

}

func testLogger(t *testing.T) *log.Logger {
	return log.New(testWriter{t}, "test", log.LstdFlags)
}

type testWriter struct {
	t *testing.T
}

func (tw testWriter) Write(p []byte) (n int, err error) {
	tw.t.Log(string(p))
	return len(p), nil
}
