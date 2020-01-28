package pagerduty

import (
	"time"
)

type Client struct{}

func (c Client) OnCallThisWeek() (string, error) {
	// imitating network latency
	time.Sleep(3 * time.Second)
	// this would change next week, making the E2E test fragile
	return "Jozsi", nil
}
