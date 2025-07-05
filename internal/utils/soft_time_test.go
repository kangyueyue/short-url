package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestSoftTime(t *testing.T) {
	timer := NewSoftTimer(1)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(timer.Time)
	}
}
