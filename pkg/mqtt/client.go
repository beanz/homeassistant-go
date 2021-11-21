package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
)

type Msg struct {
	Topic  string
	Body   interface{}
	Retain bool
}

type Client struct {
	broker *url.URL
	config *ClientConfig
	logger *log.Logger
	cm     *autopaho.ConnectionManager
}

func NewClient(cfg *ClientConfig, logger *log.Logger) (*Client, error) {
	err := ValidateConfig(cfg)
	if err != nil {
		return nil, err
	}
	brokerURL, err := url.Parse(cfg.Broker)
	if err != nil {
		return nil, fmt.Errorf("invalid broker url: %w", err)
	}
	return &Client{broker: brokerURL, config: cfg, logger: logger}, nil
}

func ValidateConfig(cfg *ClientConfig) error {
	if cfg.AppName == "" {
		cfg.AppName = filepath.Base(os.Args[0])
	}
	if cfg.Version == "" {
		cfg.Version = "0.0.0"
	}
	if cfg.ClientID == "" {
		return fmt.Errorf("ClientConfig is missing required ClientID")
	}
	if cfg.Broker == "" {
		return fmt.Errorf("ClientConfig is missing required Broker")
	}
	if cfg.KeepAlive == 0 {
		cfg.KeepAlive = 30
	}
	if cfg.ConnectRetryDelay == 0 {
		cfg.ConnectRetryDelay = time.Second * 10
	}
	if cfg.Log == nil {
		cfg.Log = log.New(os.Stdout, "ha/mqttc",
			log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	if cfg.DiscoveryTopicPrefix == "" {
		cfg.DiscoveryTopicPrefix = "test/discovery"
	} else if cfg.DiscoveryTopicPrefix[len(cfg.DiscoveryTopicPrefix)-1] == '/' {
		cfg.DiscoveryTopicPrefix = cfg.DiscoveryTopicPrefix[:len(cfg.DiscoveryTopicPrefix)-1]
	}
	if cfg.DataTopicPrefix == "" {
		cfg.DataTopicPrefix = TopicSafe(cfg.AppName)
	} else if cfg.DataTopicPrefix[len(cfg.DataTopicPrefix)-1] == '/' {
		cfg.DataTopicPrefix = cfg.DataTopicPrefix[:len(cfg.DataTopicPrefix)-1]
	}
	return nil
}

func (c *Client) Run(ctx context.Context, in chan *Msg, out chan *Msg) error {
	c.logger.Printf("%s v%s\n", c.config.AppName, c.config.Version)
	c.logger.Println("Starting Home Assistant MQTT client")

	childCtx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	bridgeAvailabilityTopic := AvailabilityTopic(c.config.DataTopicPrefix,
		"bridge")

	cmCfg := autopaho.ClientConfig{
		BrokerUrls:        []*url.URL{c.broker},
		KeepAlive:         uint16(c.config.KeepAlive),
		ConnectRetryDelay: c.config.ConnectRetryDelay,
		OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
			c.logger.Println("MQTT connection up")
			c.cm = cm
			err := c.publish(childCtx, bridgeAvailabilityTopic, "online", true)
			if err != nil {
				c.logger.Printf(
					"failed to publish bridge availability online message: %s",
					err)
			}
		},
		OnConnectError: func(err error) {
			c.logger.Printf("error whilst attempting connection: %s\n", err)
		},
		Debug: paho.NOOPLogger{},
		ClientConfig: paho.ClientConfig{
			ClientID: c.config.ClientID,
			OnClientError: func(err error) {
				c.logger.Printf("server requested disconnect: %s\n", err)
			},
			OnServerDisconnect: func(d *paho.Disconnect) {
				if d.Properties != nil {
					c.logger.Printf("server requested disconnect: %s\n",
						d.Properties.ReasonString)
				} else {
					c.logger.Printf(
						"server requested disconnect; reason code: %d\n",
						d.ReasonCode)
				}
			},
		},
	}
	c.logger.Printf("setting will message %s: %s\n",
		bridgeAvailabilityTopic, "offline")
	cmCfg.SetWillMessage(bridgeAvailabilityTopic, []byte("offline"), 1, true)

	cm, err := autopaho.NewConnection(childCtx, cmCfg)
	if err != nil {
		return err
	}
	c.cm = cm

LOOP:
	for {
		err = cm.AwaitConnection(childCtx)
		if err != nil { // Should only happen when context is cancelled
			return fmt.Errorf("broker connection error: %s", err)
		}

		select {
		case m := <-in:
			err = c.publish(ctx, m.Topic, m.Body, m.Retain)
			if err != nil {
				c.logger.Printf(
					"failed to publish message for %s: %s",
					m.Topic, err)
			}
		case <-ctx.Done():
			break LOOP
		}
	}
	c.logger.Println("shutting down")

	ctx, cancel = context.WithTimeout(childCtx, 5*time.Second)
	defer func() {
		cancel()
	}()

	pr, err := cm.Publish(ctx, &paho.Publish{
		QoS:     1,
		Topic:   bridgeAvailabilityTopic,
		Payload: []byte("offline"),
		Retain:  true,
	})
	if err != nil {
		fmt.Printf("failed to publish availability offline message: %s\n", err)
	} else if pr.ReasonCode != 0 && pr.ReasonCode != 16 {
		// 16 = Server received message but there are no subscribers
		fmt.Printf("publish availability offline reason code %dn",
			pr.ReasonCode)
	}
	_ = cm.Disconnect(ctx)

	return nil
}

func (c *Client) publish(ctx context.Context, topic string, body interface{}, retain bool) error {
	var b []byte
	var err error
	if s, ok := body.(string); ok {
		b = []byte(s)
	} else {
		b, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	go func(msg []byte) {
		pr, err := c.cm.Publish(ctx, &paho.Publish{
			QoS:     1,
			Topic:   topic,
			Payload: msg,
			Retain:  retain,
		})
		if err != nil {
			c.logger.Printf("error publishing: %s\n", err)
		} else if pr.ReasonCode != 0 && pr.ReasonCode != 16 {
			// 16 = Server received message but there are no subscribers
			c.logger.Printf("reason code %d received\n", pr.ReasonCode)
		}
	}(b)
	return nil
}
