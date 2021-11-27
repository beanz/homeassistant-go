package mqtt

import (
	"context"
	"log"
	"time"
)

type ClientConfig struct {
	AppName              string
	Version              string
	Debug                bool
	ClientID             string
	Broker               string
	KeepAlive            int16
	ConnectRetryDelay    time.Duration
	Log                  *log.Logger
	DiscoveryTopicPrefix string
	DataTopicPrefix      string
	Subs                 []Sub
}

type Msg struct {
	Topic  string
	Body   interface{}
	Retain bool
}

type Sub struct {
	Topic string
	QoS   byte
}

type PubSubServer interface {
	Run(ctx context.Context, in chan *Msg, out chan *Msg) error
}
