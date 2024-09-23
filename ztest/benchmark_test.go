package test

import (
	"testing"
	"time"
)

func Benchmark(t *testing.B) {
	logger := mockZapLogger()
	logger.WithMonitor(nil)
	logger.WithLogWrapper(nil)
	parmaMap := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
	}
	satrt := time.Now()
	ch := make(chan int, 10)
	routineNum := 1
	for j := 0; j < routineNum; j++ {
		go func() {
			for i := 0; i < 100000; i++ {
				logger.Infow("Hi every one, very nice to meet you. I am testing log framework!", parmaMap)
			}
			ch <- 1
		}()
	}
	for j := 0; j < routineNum; j++ {
		<-ch
	}
	end := time.Now()
	takenSecond := float64(end.Sub(satrt).Milliseconds()) / float64(1000)
	t.Logf("takenSecond:%f", takenSecond)
}
