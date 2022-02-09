package govmq

type QoS int64

const (
	One   QoS = 1
	Two       = 2
	Three     = 3
)

type TopicConfig struct {
	Topic string `json:"topic"`
	QoS   QoS    `json:"qos"`
}
