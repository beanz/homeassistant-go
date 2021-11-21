package mqtt

import (
	"context"
	"log"
	"time"
)

type ClientConfig struct {
	AppName              string
	Version              string
	ClientID             string
	Broker               string
	KeepAlive            int16
	ConnectRetryDelay    time.Duration
	Log                  *log.Logger
	DiscoveryTopicPrefix string
	DataTopicPrefix      string
}

type PubSubServer interface {
	Run(ctx context.Context, in chan *Msg, out chan *Msg) error
}
