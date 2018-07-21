package serve

import "time"

type Config struct {
	address string
	port    int
	timeout time.Duration
}