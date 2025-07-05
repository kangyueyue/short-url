package utils

import (
	"time"
)

// SoftTimer 软定时器
type SoftTimer struct {
	ticker *time.Ticker
	Time   int64
	stopCh chan struct{}
}

// NewSoftTimer new
func NewSoftTimer(interval time.Duration) *SoftTimer {
	s := &SoftTimer{
		ticker: time.NewTicker(time.Second * interval),
		stopCh: make(chan struct{}),
		Time:   int64(0),
	}
	go s.run()
	return s
}

// run
func (s *SoftTimer) run() {
	defer s.ticker.Stop()
	for {
		select {
		case <-s.ticker.C:
			s.Time++
		case <-s.stopCh:
			return
		}
	}
}

// Stop 停止
func (s *SoftTimer) Stop() {
	close(s.stopCh)
	s.ticker.Stop()
}
