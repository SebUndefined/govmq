package govmq

type QoS int64

const (
	Zero QoS = 0
	One      = 1
	Two      = 2
)

type TopicConfig struct {
	Topic string `json:"topic"`
	QoS   QoS    `json:"qos"`
}
